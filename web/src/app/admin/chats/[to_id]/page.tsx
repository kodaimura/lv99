import styles from "./page.module.css";
import ChatArea from "./chat-area";
import { api } from "@/lib/api/api.server";
import { AccountWithProfile } from "@/types/models";
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: "lv99 - チャット",
};

type Props = {
  params: Promise<{ to_id: string }>
};

const ChatPage: React.FC<Props> = async ({ params }) => {
  const { to_id } = await params;
  const account: AccountWithProfile = await api.get(`/admin/accounts/${to_id}/with-profile`);

  return (
    <div className={styles.page}>
      <div className={styles.header}>{account.display_name}</div>
      <div className={styles.chatAreaWrapper}>
        <ChatArea toId={parseInt(to_id)} />
      </div>
    </div>
  );
}

export default ChatPage;