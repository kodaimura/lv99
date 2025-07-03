import React from 'react';
import styles from './page.module.css';
import { AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.server';
import Link from 'next/link';

const AccountsPage: React.FC = async () => {
  const accounts: AccountWithProfile[] = await api.get('admin/accounts/with-profile');

  const formatDate = (dateStr: string | null) => {
    if (!dateStr) return null;
    const date = new Date(dateStr);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}`;
  };

  return (
    <div className={styles.container}>
      <div className={styles.tableContainer}>
        <table className={styles.table}>
          <thead className={styles.thead}>
            <tr>
              <th className={styles.th}>ID</th>
              <th className={styles.th}>アカウント名</th>
              <th className={styles.th}>表示名</th>
              <th className={styles.th}>登録日</th>
              <th className={styles.th}>回答状況</th>
            </tr>
          </thead>
          <tbody>
            {accounts.map((account) => (
              <tr key={account.id} className={styles.tr}>
                <td className={styles.td}>{account.id}</td>
                <td className={styles.td}>{account.name}</td>
                <td className={styles.td}>{account.display_name}</td>
                <td className={styles.td}>{formatDate(account.created_at)}</td>
                <td className={styles.td}>
                  <Link href={`/admin/accounts/${account.id}`}>回答状況</Link>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default AccountsPage;
