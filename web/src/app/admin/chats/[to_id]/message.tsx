'use client';

import React from "react";
import styles from "./message.module.css";
import { Chat } from "@/types/models";
import { useAuth } from "@/contexts/auth-context";

type Props = {
  chat: Chat;
};

const Message: React.FC<Props> = ({ chat }) => {
  const { account } = useAuth();
  const isMe = chat.from_id === account?.id;

  return (
    <div className={`${styles.messageRow} ${isMe ? styles.me : styles.other}`}>
      <div className={styles.bubble}>
        <div className={styles.meta}>
          <span className={styles.timestamp}>
            {new Date(chat.created_at).toLocaleTimeString()}
          </span>
        </div>
        <div className={styles.text}>{chat.message}</div>
      </div>
    </div>
  );
}

export default Message;
