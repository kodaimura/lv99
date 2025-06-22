import React from 'react';
import { AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';

type Props = {
  params: Promise<{ id: number }>
};

const AccountDetailPage: React.FC<Props> = async ({ params }) => {
  const { id } = await params
  const account: AccountWithProfile = await api.get(`admin/accounts/${id}/with-profile`);

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