import { HttpError } from './common';

class Api {
  url: string;

  constructor(url: string) {
    this.url = url;
  }

  private async createFetchOptions(method: string, body?: unknown): Promise<RequestInit> {
    const options: RequestInit = {
      method,
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    };

    if (body) {
      options.body = JSON.stringify(body);
    }
    return options;
  }

  private async apiFetch<T>(endpoint: string, method: string, body?: unknown, retry = true): Promise<T> {
    if (endpoint.startsWith('/')) {
      endpoint = endpoint.slice(1);
    }

    const options = await this.createFetchOptions(method, body);
    const response = await fetch(`${this.url}/${endpoint}`, options);

    if (!response.ok) {
      if (response.status === 401 && retry && window.location.pathname !== '/login') {
        const refreshed = await this.tryRefreshToken();
        if (refreshed) {
          return this.apiFetch<T>(endpoint, method, body, false);
        }
      }

      let errorData = { error: 'Unknown error', details: {} };
      try {
        errorData = await response.json();
      } catch {
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

  private async tryRefreshToken(): Promise<boolean> {
    try {
      const response = await fetch(`${this.url}/accounts/refresh`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      });

      return response.ok;
    } catch {
      return false;
    }
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

  async patch<T>(endpoint: string, body?: unknown): Promise<T> {
    return this.apiFetch<T>(endpoint, 'PATCH', body);
  }

  handleHttpError(error: HttpError): void {
    console.error(error);
    const status = error.status
    if (status === 403) {
      alert('アクセスが拒否されました');
    } else if (status === 401 && window.location.pathname !== '/login') {
      window.location.replace('/login');
    }
  }
}

const api = new Api(`${process.env.NEXT_PUBLIC_API_HOST}/api`);

export { Api, api };
