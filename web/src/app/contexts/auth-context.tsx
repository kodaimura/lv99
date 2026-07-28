import { Account } from '@/types/models';
import { createContext, useCallback, useContext, useEffect, useState, ReactNode } from 'react';
import { api } from '@/lib/api/api.client';

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
  try {
    return await api.get<Account>('accounts/me');
  } catch {
    return null;
  }
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
