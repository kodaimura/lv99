import { Account } from '@/types/models';
import { createContext, useCallback, useContext, useEffect, useState, ReactNode } from 'react';
import { getApiBaseUrl } from '@/lib/env';

type AuthContextType = {
  account: Account | null;
  isLoading: boolean;
  setAccount: (account: Account | null) => void;
  refreshAccount: () => Promise<Account | null>;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

type Props = {
  children: ReactNode;
  initAccount?: Account | null;
};

const fetchAccount = async (): Promise<Account | null> => {
  const baseUrl = getApiBaseUrl();

  const requestMe = () => fetch(`${baseUrl}/accounts/me`, {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
  });

  let response = await requestMe();
  if (response.status === 401) {
    const refresh = await fetch(`${baseUrl}/accounts/refresh`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    });
    if (refresh.ok) response = await requestMe();
  }

  if (!response.ok) return null;
  return (await response.json()) as Account;
};

export const AuthProvider: React.FC<Props> = ({ children, initAccount = null }) => {
  const [account, setAccount] = useState<Account | null>(initAccount);
  const [isLoading, setIsLoading] = useState(true);

  const refreshAccount = useCallback(async () => {
    try {
      const nextAccount = await fetchAccount();
      setAccount(nextAccount);
      return nextAccount;
    } catch {
      setAccount(null);
      return null;
    }
  }, []);

  useEffect(() => {
    let active = true;

    const hydrate = async () => {
      try {
        const nextAccount = await fetchAccount();
        if (active) setAccount(nextAccount);
      } catch {
        if (active) setAccount(null);
      } finally {
        if (active) setIsLoading(false);
      }
    };

    hydrate();
    return () => {
      active = false;
    };
  }, []);

  return (
    <AuthContext.Provider value={{ account, isLoading, setAccount, refreshAccount }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
