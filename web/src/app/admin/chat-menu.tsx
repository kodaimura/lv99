'use client';

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import styles from './chat-menu.module.css';
import { AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.client';
import { useAuth } from '@/contexts/auth-context';

type UnreadCountMap = {
  [accountId: string]: {
    count: number;
    updatedAt: string;
  };
};

const ChatMenu: React.FC = () => {
  const router = useRouter();
  const [accounts, setAccounts] = useState<AccountWithProfile[]>([]);
  const [unreadCounts, setUnreadCounts] = useState<UnreadCountMap>({});

  const { account } = useAuth();

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
      if (data.from_id === account?.id) return;
      setUnreadCounts((prev) => ({
        ...prev,
        [data.from_id]: {
          count: (prev[data.from_id]?.count || 0) + 1,
          updatedAt: new Date().toISOString(),
        },
      }));

      const audio = new Audio('/message.mp3');
      audio.play();
    };

    socket.onclose = () => {
      console.log("WebSocket closed for chat menu");
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

  const handleClickChat = (accountId: number) => {
    setUnreadCounts((prev) => ({
      ...prev,
      [accountId]: { ...prev[accountId], count: 0, },
    }));
    router.push(`/admin/chats/${accountId}`);
  }

  return (
    <div className={styles.chatMenu}>
      {sortedAccounts.map((account) => (
        <a
          key={account.id}
          className={styles.item}
          onClick={(e) => {
            e.preventDefault();
            handleClickChat(account.id);
          }}
        >
          <span className={styles.name}>{account.display_name}</span>
          {unreadCounts[account.id]?.count > 0 && (
            <span className={styles.badge}>
              {unreadCounts[account.id].count}
            </span>
          )}
        </a>
      ))}
    </div>
  );
};

export default ChatMenu;
