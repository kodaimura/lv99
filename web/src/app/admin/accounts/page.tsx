'use client';

import React, { useEffect, useState } from 'react';
import styles from './page.module.css';
import { AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.client';
import Link from 'next/link';
import Modal from '@/components/ui/modal';
import AccountForm from './account-form';
import LocalDate from '@/components/features/local-date';

const AccountsPage: React.FC = () => {
  const [showModal, setShowModal] = useState(false);
  const [accounts, setAccounts] = useState<AccountWithProfile[]>([]);

  const getAccounts = async () => {
    try {
      const response: AccountWithProfile[] = await api.get('admin/accounts/with-profile');
      setAccounts(response);
    } catch (error) {
      console.error('Failed to fetch accounts:', error);
    }
  };

  const handleSuccess = () => {
    setShowModal(false);
    getAccounts();
  };

  useEffect(() => {
    getAccounts();
  }, [])

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <button onClick={() => setShowModal(true)} className={styles.addButton}>
          アカウント追加
        </button>
      </div>
      <Modal isOpen={showModal} onClose={() => setShowModal(false)} title='アカウント追加' >
        <AccountForm onSuccess={handleSuccess} />
      </Modal>
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
                <td className={styles.td}><LocalDate isoString={account.created_at} /></td>
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
