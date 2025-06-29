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

  const formatDate = (dateStr: string) => {
    const date = new Date(dateStr);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}`;
  };

  return (
    <div className={`${styles.messageRow} ${isMe ? styles.me : styles.other}`}>
      <div className={styles.bubble}>
        <div className={styles.meta}>
          <span className={styles.timestamp}>
            {formatDate(chat.created_at)}
          </span>
        </div>
        <div className={styles.text}>{chat.message}</div>
      </div>
    </div>
  );
}

export default Message;
