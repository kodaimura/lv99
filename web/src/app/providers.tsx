'use client';

import { AuthProvider } from '@/contexts/auth-context';
import { Account } from '@/types/models';

export const Providers = ({
  children,
  initAccount,
}: {
  children: React.ReactNode;
  initAccount: Account | null;
}) => {
  return <AuthProvider initAccount={initAccount}>{children}</AuthProvider>;
};
