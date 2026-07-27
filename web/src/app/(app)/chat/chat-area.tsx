'use client';

import { useCallback, useEffect, useRef, useState } from "react";
import styles from "./chat-area.module.css";
import Message from "./message";
import { Chat } from "@/types/models";
import { api } from "@/lib/api/api.client";
import { useReliableWebSocket } from '@/hooks/websocket';

type Props = {
  toId: number;
};

const parseChatTime = (chat: Chat) => {
  const time = Date.parse(chat.created_at);
  return Number.isNaN(time) ? 0 : time;
};

const sortChats = (items: Chat[]) =>
  items.sort((a, b) => parseChatTime(a) - parseChatTime(b) || a.id - b.id);

const mergeChats = (current: Chat[], incoming: Chat[]) => {
  const byId = new Map<number, Chat>();
  for (const chat of current) byId.set(chat.id, chat);
  for (const chat of incoming) byId.set(chat.id, chat);
  return sortChats(Array.from(byId.values()));
};

const ChatArea: React.FC<Props> = ({ toId }) => {
  const [chats, setChats] = useState<Chat[]>([]);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const bottomRef = useRef<HTMLDivElement>(null);
  const chatBoxRef = useRef<HTMLDivElement>(null);
  const chatsRef = useRef<Chat[]>([]);
  const loadingRef = useRef(false);
  const requestSeqRef = useRef(0);

  const { ready, send, onMessage } = useReliableWebSocket(
    `${process.env.NEXT_PUBLIC_WS_HOST}/ws/chats`
  );
  const [message, setMessage] = useState("");

  useEffect(() => {
    chatsRef.current = chats;
  }, [chats]);

  const scrollToBottom = useCallback(() => {
    requestAnimationFrame(() => {
      const box = chatBoxRef.current;
      if (box) {
        box.scrollTop = box.scrollHeight;
      }
    });
  }, []);

  const getChats = useCallback(async (isInit = false) => {
    if (loadingRef.current) return;
    loadingRef.current = true;
    setLoading(true);

    const requestSeq = requestSeqRef.current;
    const oldest = isInit ? undefined : chatsRef.current[0];
    const query = oldest ? `?before=${encodeURIComponent(oldest.created_at)}` : "";

    try {
      const response: Chat[] = await api.get(`chats/${toId}${query}`);
      if (requestSeq !== requestSeqRef.current) return;

      if (response.length === 0) {
        setHasMore(false);
        return;
      }

      const normalized = response.slice().reverse();
      setChats((prev) => mergeChats(isInit ? [] : prev, normalized));
      scrollToBottom();
    } catch (e) {
      console.error(e);
    } finally {
      if (requestSeq === requestSeqRef.current) {
        loadingRef.current = false;
        setLoading(false);
      }
    }
  }, [scrollToBottom, toId]);

  const readChats = useCallback(async () => {
    try {
      await api.put("chats/read", { from_id: toId });
    } catch (e) {
      console.error("Failed to mark chats as read:", e);
    }
  }, [toId]);

  useEffect(() => {
    requestSeqRef.current += 1;
    loadingRef.current = false;
    setLoading(false);
    setChats([]);
    chatsRef.current = [];
    setHasMore(true);
    getChats(true);
    readChats();
    requestAnimationFrame(() => {
      const box = chatBoxRef.current;
      if (box) box.scrollTop = box.scrollHeight;
    });
  }, [getChats, readChats]);

  useEffect(() => {
    const off = onMessage((event) => {
      const data = JSON.parse(event.data);
      if (data.from_id !== toId && data.to_id !== toId) return;
      setChats((prev) => mergeChats(prev, [data]));
      scrollToBottom();
    });
    return off;
  }, [toId, onMessage, scrollToBottom]);

  const sendMessage = () => {
    const text = message.trim();
    if (!text) return;
    send(JSON.stringify({ to_id: Number(toId), message: text }));
    setMessage("");
  };

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
  }, [getChats, hasMore, loading]);

  return (
    <div className={styles.container}>
      <div className={styles.chatBox} ref={chatBoxRef}>
        {chats.map((chat) => (
          <Message key={chat.id} chat={chat} />
        ))}
        <div ref={bottomRef} />
      </div>
      <div className={styles.inputArea}>
        <textarea
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          className={styles.input}
          placeholder="Type a message..."
          rows={3}
        />
        <button
          onClick={sendMessage}
          className={styles.button}
          disabled={!ready || !message.trim()}
        >
          送信
        </button>
      </div>
    </div>
  );
}

export default ChatArea;
