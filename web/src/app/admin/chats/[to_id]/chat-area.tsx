'use client';

import { useEffect, useState } from "react";
import styles from "./chat-area.module.css";
import Message from "./message";
import { Chat } from "@/types/models";
import { api } from "@/lib/api/api.client";
import ChatForm from "./chat-form";

type Props = {
  toId: number;
};

const ChatArea: React.FC<Props> = ({ toId }) => {
  const [chats, setChats] = useState<Chat[]>([]);

  const getChats = async () => {
    try {
      const response: Chat[] = await api.get(`chats/${toId}`);
      setChats(response);
    } catch (e) {
      console.error(e)
    }
  }

  const handleRecieve = (chat: Chat) => {
    console.log(chat)
    setChats(prev => [...prev, chat]);
  }

  useEffect(() => {
    getChats();
  }, []);

  useEffect(() => {
    console.log("chats updated:", chats);
  }, [chats]);

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Chat</h1>
      <div className={styles.chatBox}>
        {chats.map((chat, i) => (
          <Message key={i} chat={chat} />
        ))}
      </div>
      <div className={styles.inputArea}>
        <ChatForm toId={toId} onRecieve={handleRecieve} />
      </div>
    </div>
  );
}

export default ChatArea;