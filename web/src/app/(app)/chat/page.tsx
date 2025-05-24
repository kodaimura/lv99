"use client";

import { useEffect, useRef, useState } from "react";
import styles from "./page.module.css";
import { useAuth } from "@/contexts/auth-context";

const ChatPage = () => {
  const { account } = useAuth();
  const [messages, setMessages] = useState<string[]>([]);
  const [message, setMessage] = useState("");
  const socketRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8000/api/chats/ws");
    socketRef.current = socket;

    socket.onmessage = (event) => {
      setMessages((prev) => [...prev, event.data]);
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
        to_id: 1,
        message: message,
      }));
      setMessage("");
    }
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Chat</h1>
      <div className={styles.chatBox}>
        {messages.map((msg, idx) => (
          <div key={idx} className={styles.message}>
            {msg}
          </div>
        ))}
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
          Send
        </button>
      </div>
    </div>
  );
}

export default ChatPage;