'use client';

import React, { use, useEffect } from 'react';
import styles from './chat-menu.module.css';
import { Account } from '@/types/models';
import { api } from '@/lib/api/api.client';
import Link from 'next/link';

const ChatMenu: React.FC = () => {
  const [accounts, setAccounts] = React.useState<Account[]>([]);

  const getAccounts = async () => {
    try {
      const response: Account[] = await api.get('admin/accounts')
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
          {account.name}
        </Link>
      ))}
    </div>
  );
};

export default ChatMenu;
