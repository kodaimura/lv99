'use client';

import React, { use, useEffect } from 'react';
import styles from './chat-menu.module.css';
import { AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.client';
import Link from 'next/link';

const ChatMenu: React.FC = () => {
  const [accounts, setAccounts] = React.useState<AccountWithProfile[]>([]);

  const getAccounts = async () => {
    try {
      const response: AccountWithProfile[] = await api.get('admin/accounts/with-profile')
      setAccounts(response);
    } catch (error) {
      console.error('Failed to fetch accounts:', error);
    }
  };

  useEffect(() => {
    getAccounts();
  }, []);

  return (
    <div className={styles.chatMenu}>
      {accounts.map((account) => (
        <Link
          key={account.id}
          className={styles.item}
          href={`/admin/chats/${account.id}`}
        >
          {account.display_name}
        </Link>
      ))}
    </div>
  );
};

export default ChatMenu;
