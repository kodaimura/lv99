let accessToken: string | null = null;

export const getAccessToken = () => accessToken;

export const setAccessToken = (token: string | null | undefined) => {
  accessToken = token || null;
};

export const clearAccessToken = () => {
  accessToken = null;
};

export const getWebSocketAuthProtocols = () => {
  if (!accessToken) return null;
  return ['lv99', `access_token.${accessToken}`];
};
