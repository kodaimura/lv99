import React from 'react';
import styles from './account-list.module.css';
import type { AccountWithProfile } from "@/types/models";
import Link from 'next/link';

type Props = {
  accounts: AccountWithProfile[];
};

const AccountList: React.FC<Props> = ({
  accounts,
}) => {

  const formatDate = (dateStr: string) => {
    const date = new Date(dateStr);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}`;
  };

  return (
    <div className={styles.tableContainer}>
      <table className={styles.table}>
        <thead className={styles.thead}>
          <tr>
            <th className={styles.th}>#</th>
            <th className={styles.th}>ユーザー名</th>
            <th className={styles.th}>作成日</th>
          </tr>
        </thead>
        <tbody className={styles.tbody}>
          {accounts.map((a, i) => (
            <tr key={i} className={styles.tr}>
              <td className={styles.td}><Link href={`accounts/${a.id}`}>{a.id}</Link></td>
              <td className={styles.td}>{a.display_name}</td>
              <td className={styles.td}>{formatDate(a.created_at)}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default AccountList;
