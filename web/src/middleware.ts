import { NextResponse, NextRequest } from 'next/server';

const isAdminPath = (pathname: string) => pathname.startsWith('/admin');

const isAdminAccount = async (access_token: string): Promise<boolean> => {
  try {
    const res = await fetch(`${process.env.API_HOST}/api/accounts/me`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Cookie: `access_token=${access_token}`,
      },
      credentials: 'include',
    });
    return (await res.json())?.role === 0
  } catch {
    return false;
  }
}

export async function middleware(request: NextRequest) {
  const accessToken = request.cookies.get('access_token')?.value;
  const accessTokenExpiresAt = request.cookies.get('access_token_expires_at')?.value;
  const refreshToken = request.cookies.get('refresh_token')?.value;

  const now = Math.floor(Date.now() / 1000);
  const bufferSeconds = 30;

  if (accessToken && accessTokenExpiresAt && now < parseInt(accessTokenExpiresAt) - bufferSeconds) {
    if (isAdminPath(request.nextUrl.pathname) && !(await isAdminAccount(accessToken))) {
      return NextResponse.redirect(new URL('/login', request.url));
    }
    return NextResponse.next();
  }

  if (!refreshToken) {
    return NextResponse.redirect(new URL('/login', request.url));
  }

  // refresh
  const res = await fetch(`${process.env.API_HOST}/api/accounts/refresh`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Cookie: `refresh_token=${refreshToken}`,
    },
    credentials: 'include',
  });

  if (res.status === 401) {
    return NextResponse.redirect(new URL('/login', request.url));
  }

  const { access_token, expires_in } = await res.json();
  const headerSetCookies = res.headers.get('set-cookie');
  const response = NextResponse.next();

  if (access_token) {
    response.cookies.set({
      name: 'access_token',
      value: access_token,
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'lax',
      path: '/',
    });
  }

  // Append Set-Cookie: access_token_expires_at
  if (expires_in) {
    response.headers.append('Set-Cookie', [
      `access_token_expires_at=${(now + expires_in).toString()}`,
      'HttpOnly',
      process.env.NODE_ENV === 'production' ? 'Secure' : '',
      'SameSite=Lax',
      'Path=/',
    ].filter(Boolean).join('; '));
  }

  // Append other Set-Cookie headers from API response
  if (headerSetCookies) {
    const cookies = headerSetCookies.split(',');
    cookies.forEach(cookie => response.headers.append('Set-Cookie', cookie.trim()));
  }

  if (isAdminPath(request.nextUrl.pathname) && !(await isAdminAccount(access_token))) {
    return NextResponse.redirect(new URL('/login', request.url));
  }

  return response;
}

export const config = {
  matcher: [
    '/dashboard',
    '/admin/:path*',
    '/questions/:path*',
    '/chat/:path*',
    '/account',
  ],
};