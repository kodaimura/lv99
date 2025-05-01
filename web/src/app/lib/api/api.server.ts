import { cookies } from 'next/headers';
import { redirect, forbidden } from 'next/navigation';
import { HttpError } from './common';

class Api {
  url: string;

  constructor(url: string) {
    this.url = url;
  }

  public async createFetchOptions(method: string, body?: unknown): Promise<RequestInit> {
    const cookieStore = await cookies();
    const accessToken = cookieStore.get('access_token')?.value ?? "";

    const options: RequestInit = {
      method,
      headers: {
        'Content-Type': 'application/json',
        Cookie: `access_token=${accessToken}`,
      },
      credentials: 'include',
    };

    if (body) {
      options.body = JSON.stringify(body);
    }

    return options;
  }

  private async apiFetch<T>(endpoint: string, method: string, body?: unknown): Promise<T> {
    if (endpoint.startsWith('/')) {
      endpoint = endpoint.slice(1);
    }

    const options = await this.createFetchOptions(method, body);
    const response = await fetch(`${this.url}/${endpoint}`, options);

    if (!response.ok) {
      let errorData = { error: 'Unknown error', details: {} };
      try {
        errorData = await response.json();
      } catch (_) {
        // ignore parse error
      }

      const error = new HttpError(response.status, errorData.error, errorData.details);
      this.handleHttpError(error);
      throw error;
    }

    if (response.status === 204) {
      return {} as T;
    }

    return (await response.json()) as T;
  }

  async get<T>(endpoint: string): Promise<T> {
    return this.apiFetch<T>(endpoint, 'GET');
  }

  async post<T>(endpoint: string, body: unknown): Promise<T> {
    return this.apiFetch<T>(endpoint, 'POST', body);
  }

  async put<T>(endpoint: string, body: unknown): Promise<T> {
    return this.apiFetch<T>(endpoint, 'PUT', body);
  }

  async delete<T>(endpoint: string, body?: unknown): Promise<T> {
    return this.apiFetch<T>(endpoint, 'DELETE', body);
  }

  public handleHttpError(error: HttpError): void {
    console.error(error)
    const status = error.status;
    if (status === 401) {
      redirect('/login');
    } else if (status === 403) {
      forbidden();
    }
  }
}

const api = new Api(`${process.env.API_HOST}/api`);

export { Api, api };
