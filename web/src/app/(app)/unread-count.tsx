'use client';

import React, { useEffect, useState } from 'react';
import { AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.client';
import styles from './unread-count.module.css';
import { useAuth } from '@/contexts/auth-context';

const UnreadCount: React.FC = () => {
  const [unreadCount, setUnreadCount] = useState(0);
  const [admin, setAdmin] = useState<AccountWithProfile | null>(null);
  const { account } = useAuth();

  useEffect(() => {
    const fetchAdmin = async () => {
      try {
        const adminData: AccountWithProfile = await api.get('/accounts/admin/with-profile');
        setAdmin(adminData);
      } catch (error) {
        console.error('Failed to fetch admin data:', error);
      }
    };
    fetchAdmin();
  }, []);

  useEffect(() => {
    if (!admin || !account) return;

    const fetchUnreadCount = async () => {
      try {
        const response: { account_id: number; unread_count: number; updated_at: string }[] =
          await api.get('chats/unread-count');
        const item = response.find((item) => item.account_id === admin.id);
        if (item) {
          setUnreadCount(item.unread_count);
        }
      } catch (error) {
        console.error('Failed to fetch unread count:', error);
      }
    };

    fetchUnreadCount();

    const socket = new WebSocket("ws://localhost:8000/api/chats/ws");

    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.from_id === account.id) return;
      if (data.to_id !== admin.id) return;
      setUnreadCount((prev) => prev + 1);
    };

    socket.onclose = () => {
      console.log("WebSocket closed for chat menu");
    };

    return () => {
      socket.close();
    };
  }, [admin, account]);

  return (
    <>
      {unreadCount > 0 && (
        <span className={styles.badge}>
          {unreadCount}
        </span>
      )}
    </>
  );
};

export default UnreadCount;
