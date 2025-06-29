'use client';

import { useEffect, useRef, useState } from "react";
import styles from "./chat-area.module.css";
import Message from "./message";
import { Chat } from "@/types/models";
import { api } from "@/lib/api/api.client";

type Props = {
  toId: number;
};

const ChatArea: React.FC<Props> = ({ toId }) => {
  const [chats, setChats] = useState<Chat[]>([]);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const bottomRef = useRef<HTMLDivElement>(null);
  const chatBoxRef = useRef<HTMLDivElement>(null);

  const socketRef = useRef<WebSocket | null>(null);
  const [message, setMessage] = useState("");

  useEffect(() => {
    getChats(true);
    readChats();
    const socket = new WebSocket("ws://localhost:8000/api/chats/ws");
    socketRef.current = socket;

    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.from_id !== toId && data.to_id !== toId) return;
      setChats(prev => [...prev, data]);
      scrollToBottom();
    };

    socket.onclose = () => {
      console.log("WebSocket closed");
    };

    return () => {
      socket.close();
    };
  }, []);

  const sendMessage = () => {
    if (socketRef.current && message.trim()) {
      socketRef.current.send(JSON.stringify({
        to_id: Number(toId),
        message: message,
      }));
      setMessage("");
    }
  };

  const scrollToBottom = () => {
    requestAnimationFrame(() => {
      const box = chatBoxRef.current;
      if (box) {
        box.scrollTop = box.scrollHeight;
      }
    });
  }

  const getChats = async (isInit = false) => {
    if (loading) return;
    setLoading(true);

    const oldest = chats[0];
    const query = oldest ? `before=${oldest.created_at}` : "";

    try {
      const response: Chat[] = await api.get(`chats/${toId}?${query}`);
      if (response.length === 0) {
        setHasMore(false);
      } else {
        if (isInit) {
          setChats(response.slice().reverse());
        } else {
          setChats(prev => [...response.slice().reverse(), ...prev]);
        }
        scrollToBottom();
      }
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  };

  const readChats = async () => {
    try {
      await api.put("chats/read", { from_id: toId });
    } catch (e) {
      console.error("Failed to mark chats as read:", e);
    }
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
  }, [hasMore, loading]);

  return (
    <div className={styles.container}>
      <div className={styles.chatBox} ref={chatBoxRef}>
        {chats.map((chat, i) => (
          <Message key={i} chat={chat} />
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
        <button onClick={sendMessage} className={styles.button}>
          送信
        </button>
      </div>
    </div>
  );
}

export default ChatArea;