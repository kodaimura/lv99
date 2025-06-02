'use client';

import { useEffect, useRef, useState } from "react";
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
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const bottomRef = useRef<HTMLDivElement>(null);
  const chatBoxRef = useRef<HTMLDivElement>(null);

  const scrollToBottom = () => {
    requestAnimationFrame(() => {
      const box = chatBoxRef.current;
      if (box) {
        box.scrollTop = box.scrollHeight;
      }
    });
  }

  const getChats = async () => {
    if (loading) return;
    setLoading(true);

    const oldest = chats[0];
    const query = oldest ? `before=${oldest.created_at}` : "";

    try {
      const response: Chat[] = await api.get(`chats/${toId}?${query}`);
      if (response.length === 0) {
        setHasMore(false);
      } else {
        setChats(prev => [...response.slice().reverse(), ...prev]);
        scrollToBottom();
      }
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  };

  const handleRecieve = (chat: Chat) => {
    setChats(prev => [...prev, chat]);
    scrollToBottom();
  }

  useEffect(() => {
    getChats();
  }, []);

  useEffect(() => {
    if (chats.length <= 30) {
      bottomRef.current?.scrollIntoView({ behavior: "auto" });
    }
  }, [chats.length]);

  useEffect(() => {
    const box = chatBoxRef.current;
    if (!box) return;

    const handleScroll = () => {
      if (box.scrollTop === 0 && hasMore && !loading) {
        getChats();
      }
    };

    box.addEventListener("scroll", handleScroll);
    return () => box.removeEventListener("scroll", handleScroll);
  }, [hasMore, loading]);

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Chat</h1>
      <div className={styles.chatBox} ref={chatBoxRef}>
        {chats.map((chat, i) => (
          <Message key={i} chat={chat} />
        ))}
        <div ref={bottomRef} />
      </div>
      <div className={styles.inputArea}>
        <ChatForm toId={toId} onRecieve={handleRecieve} />
      </div>
    </div>
  );
}

export default ChatArea;