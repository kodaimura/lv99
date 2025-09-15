'use client';

import React, { useEffect, useState, useRef } from 'react';
import { useRouter } from 'next/navigation';
import { AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.client';
import styles from './chat-menu.module.css';
import { useAuth } from '@/contexts/auth-context';

type Props = {
  onClick?: () => void;
};

const ChatMenu: React.FC<Props> = ({ onClick }) => {
  const router = useRouter();
  const [unreadCount, setUnreadCount] = useState(0);
  const [admin, setAdmin] = useState<AccountWithProfile | null>(null);
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

    const socket = new WebSocket(`${process.env.NEXT_PUBLIC_WS_HOST}/ws/chats`);

    socket.onmessage = async (event) => {
      const data = JSON.parse(event.data);
      if (data.from_id === account.id) return;
      if (data.from_id !== admin.id) return;
      setUnreadCount((prev) => prev + 1);

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
  }, [admin, account]);

  const handleClickChat = (e: React.MouseEvent<HTMLAnchorElement>) => {
    e.preventDefault();
    setUnreadCount(0);
    router.push('/chat');
    onClick?.();
  };

  return (
    <>
      <a className={styles.item} onClick={handleClickChat}>
        チャット {unreadCount > 0 && (<span className={styles.badge}>{unreadCount}</span>)}
      </a>
      <audio ref={audioRef} src="/message.mp3" preload="auto" />
    </>
  );
};

export default ChatMenu;
