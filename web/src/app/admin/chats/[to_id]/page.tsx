import { useEffect, useMemo, useState } from 'react';
import styles from './page.module.css';
import ChatArea from './chat-area';
import { api } from '@/lib/api/api.client';
import { AccountWithProfile } from '@/types/models';
import { useParams } from 'react-router-dom';

const ChatPage: React.FC = () => {
  const { to_id } = useParams<{ to_id: string }>();
  const toId = useMemo(() => Number.parseInt(to_id ?? '', 10), [to_id]);
  const [account, setAccount] = useState<AccountWithProfile | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;

    const fetchAccount = async () => {
      if (!Number.isInteger(toId)) {
        setLoading(false);
        return;
      }

      setLoading(true);
      try {
        const data: AccountWithProfile = await api.get(`/admin/accounts/${toId}/with-profile`);
        if (active) setAccount(data);
      } catch (error) {
        console.error('Failed to fetch chat account:', error);
      } finally {
        if (active) setLoading(false);
      }
    };

    fetchAccount();
    return () => {
      active = false;
    };
  }, [toId]);

  if (loading) return <div className={styles.page}>読み込み中...</div>;
  if (!account || !Number.isInteger(toId)) return <div className={styles.page}>チャット相手を取得できませんでした。</div>;

  return (
    <div className={styles.page}>
      <div className={styles.header}>{account.display_name}</div>
      <div className={styles.chatAreaWrapper}>
        <ChatArea toId={toId} />
      </div>
    </div>
  );
};

export default ChatPage;
