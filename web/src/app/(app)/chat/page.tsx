import styles from "./page.module.css";
import ChatArea from "./chat-area";
import { api } from "@/lib/api/api.server";
import { AccountWithProfile } from "@/types/models";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "lv99 - チャット",
};

const ChatPage: React.FC = async () => {
  const account: AccountWithProfile = await api.get("/accounts/admin/with-profile");

  return (
    <div className={styles.page}>
      <div className={styles.header}>{account.display_name}</div>
      <div className={styles.chatAreaWrapper}>
        <ChatArea toId={account.id} />
      </div>
    </div>
  );
}

export default ChatPage;