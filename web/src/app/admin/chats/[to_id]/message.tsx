'use client';

import React from "react";
import styles from "./message.module.css";
import { Chat } from "@/types/models";
import { useAuth } from "@/contexts/auth-context";
import LocalDate from "@/components/features/local-date";

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
            <LocalDate isoString={chat.created_at} />
          </span>
        </div>
        <div className={styles.text}>{chat.message}</div>
      </div>
    </div>
  );
}

export default Message;
