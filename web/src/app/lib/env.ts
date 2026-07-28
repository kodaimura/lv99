const trimTrailingSlash = (value: string) => value.replace(/\/+$/, '');

export const getApiBaseUrl = () => {
  const host = trimTrailingSlash(import.meta.env.VITE_API_HOST ?? '');
  return host ? `${host}/api` : '/api';
};

export const getWsBaseUrl = () => {
  const configured = trimTrailingSlash(import.meta.env.VITE_WS_HOST ?? '');
  if (configured) return configured;

  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
  return `${protocol}//${window.location.host}`;
};
