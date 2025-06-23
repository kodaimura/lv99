import React from 'react';
import { AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';

type Props = {
  params: { id: string }
};

const AccountDetailPage: React.FC<Props> = async ({ params }) => {
  const { id } = params
  const account: AccountWithProfile = await api.get(`admin/accounts/${id}/with-profile`);
  const answers: AccountWithProfile = await api.get(`admin/answers/${id}/with-profile`);

  console.log(account);
  return (
    <div className={styles.container}>
      <div className={styles.header}>
      </div>

      <div className={styles.content}>
      </div>
    </div>
  );
};

export default AccountDetailPage;