import { useEffect, useState } from 'react';
import styles from './page.module.css';
import ChatArea from './chat-area';
import { api } from '@/lib/api/api.client';
import { AccountWithProfile } from '@/types/models';

const ChatPage: React.FC = () => {
  const [account, setAccount] = useState<AccountWithProfile | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;

    const fetchAdmin = async () => {
      try {
        const data: AccountWithProfile = await api.get('/accounts/admin/with-profile');
        if (active) setAccount(data);
      } catch (error) {
        console.error('Failed to fetch chat account:', error);
      } finally {
        if (active) setLoading(false);
      }
    };

    fetchAdmin();
    return () => {
      active = false;
    };
  }, []);

  if (loading) return <div className={styles.page}>読み込み中...</div>;
  if (!account) return <div className={styles.page}>チャット相手を取得できませんでした。</div>;

  return (
    <div className={styles.page}>
      <div className={styles.header}>{account.display_name}</div>
      <div className={styles.chatAreaWrapper}>
        <ChatArea toId={account.id} />
      </div>
    </div>
  );
};

export default ChatPage;
