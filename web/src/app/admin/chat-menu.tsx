'use client';

import React, { useEffect, useState, useRef } from 'react';
import { useRouter } from 'next/navigation';
import styles from './chat-menu.module.css';
import { AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.client';
import { useAuth } from '@/contexts/auth-context';

type Props = {
  onClick?: () => void;
};

type UnreadCountMap = {
  [accountId: string]: {
    count: number;
    updatedAt: string;
  };
};

const ChatMenu: React.FC<Props> = ({ onClick }) => {
  const router = useRouter();
  const [accounts, setAccounts] = useState<AccountWithProfile[]>([]);
  const [unreadCounts, setUnreadCounts] = useState<UnreadCountMap>({});
  const { account } = useAuth();

  const audioRef = useRef<HTMLAudioElement | null>(null);
  const [soundEnabled, setSoundEnabled] = useState(false);

  // 初回ユーザー操作で音声を解錠する
  useEffect(() => {
    const unlock = async () => {
      const el = audioRef.current;
      if (!el) return;
      try {
        el.muted = true;
        await el.play();
        el.pause();
        el.currentTime = 0;
        el.muted = false;
        setSoundEnabled(true);
      } catch (error) {
        console.error('Failed to unlock audio:', error);
      }
    };

    // 一度だけ拾えれば良い（端末差を吸収するため3種登録）
    window.addEventListener('pointerdown', unlock, { once: true });
    window.addEventListener('keydown', unlock, { once: true });
    window.addEventListener('touchstart', unlock, { once: true });

    return () => {
      window.removeEventListener('pointerdown', unlock);
      window.removeEventListener('keydown', unlock);
      window.removeEventListener('touchstart', unlock);
    };
  }, []);

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
    if (!account?.id) return;

    getAccounts();
    getUnreadCounts();

    const socket = new WebSocket(`${process.env.NEXT_PUBLIC_WS_HOST}/ws/chats`);

    socket.onmessage = async (event) => {
      const data = JSON.parse(event.data);
      if (data.from_id === account.id) return;

      setUnreadCounts((prev) => ({
        ...prev,
        [data.from_id]: {
          count: (prev[data.from_id]?.count || 0) + 1,
          updatedAt: new Date().toISOString(),
        },
      }));

      // 解錠済みなら既存の audio を再生
      const el = audioRef.current;
      if (soundEnabled && el) {
        try {
          el.currentTime = 0;
          await el.play();
        } catch (e) {
          console.warn('play failed:', e);
        }
      }
    };

    socket.onclose = () => {
      console.log("WebSocket closed for chat menu");
    };

    return () => {
      socket.close();
    };
  }, [account?.id]);

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
    onClick?.();
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
      <audio ref={audioRef} src="/message.mp3" preload="auto" />
    </div>
  );
};

export default ChatMenu;
