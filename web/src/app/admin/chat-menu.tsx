'use client';

import React, { useEffect, useState } from 'react';
import styles from './chat-menu.module.css';
import { AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.client';
import Link from 'next/link';

type UnreadCountMap = {
  [accountId: string]: {
    count: number;
    updatedAt: string;
  };
};

const ChatMenu: React.FC = () => {
  const [accounts, setAccounts] = useState<AccountWithProfile[]>([]);
  const [unreadCounts, setUnreadCounts] = useState<UnreadCountMap>({});

  const getAccounts = async () => {
    try {
      const response: AccountWithProfile[] = await api.get('admin/accounts/with-profile');
      setAccounts(response);
    } catch (error) {
      console.error('Failed to fetch accounts:', error);
    }
  };

  const getUnreadCounts = async () => {
    try {
      const response: { account_id: number; unread_count: number; updated_at: string }[] =
        await api.get('chats/unread-count');

      const map: UnreadCountMap = {};
      response.forEach((item) => {
        map[item.account_id] = {
          count: item.unread_count,
          updatedAt: item.updated_at,
        };
      });

      setUnreadCounts(map);
    } catch (error) {
      console.error('Failed to fetch unread counts:', error);
    }
  };

  useEffect(() => {
    getAccounts();
    getUnreadCounts();

    const socket = new WebSocket("ws://localhost:8000/api/chats/ws");

    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setUnreadCounts((prev) => ({
        ...prev,
        [data.from_id]: {
          count: (prev[data.from_id]?.count || 0) + 1,
          updatedAt: new Date().toISOString(),
        },
      }));
    };

    return () => {
      socket.close();
    };
  }, []);

  const sortedAccounts = [...accounts].sort((a, b) => {
    const updatedA = unreadCounts[a.id]?.updatedAt ?? '1970-01-01T00:00:00Z';
    const updatedB = unreadCounts[b.id]?.updatedAt ?? '1970-01-01T00:00:00Z';
    return new Date(updatedB).getTime() - new Date(updatedA).getTime();
  });

  return (
    <div className={styles.chatMenu}>
      {sortedAccounts.map((account) => (
        <Link
          key={account.id}
          className={styles.item}
          href={`/admin/chats/${account.id}`}
        >
          <span className={styles.name}>{account.display_name}</span>
          {unreadCounts[account.id]?.count > 0 && (
            <span className={styles.badge}>
              {unreadCounts[account.id].count}
            </span>
          )}
        </Link>
      ))}
    </div>
  );
};

export default ChatMenu;
