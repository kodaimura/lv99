import { Account } from '@/types/models';

export const getAuthenticatedEntryPath = (account: Pick<Account, 'role'>) => {
  return account.role === 0 ? '/admin' : '/home';
};
