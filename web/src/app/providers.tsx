import { AuthProvider } from '@/contexts/auth-context';
import { Account } from '@/types/models';

export const Providers = ({
  children,
  initAccount = null,
}: {
  children: React.ReactNode;
  initAccount?: Account | null;
}) => {
  return <AuthProvider initAccount={initAccount}>{children}</AuthProvider>;
};
