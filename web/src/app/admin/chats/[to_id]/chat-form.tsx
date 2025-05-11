'use client';

import { useEffect, useRef, useState } from "react";
import styles from "./chat-form.module.css";

type Props = {
  toId: number;
};

const ChatForm: React.FC<Props> = ({ toId }) => {
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
        to_id: Number(toId),
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
        <input
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          onKeyDown={(e) => e.key === "Enter" && sendMessage()}
          className={styles.input}
          placeholder="Type a message..."
        />
        <button onClick={sendMessage} className={styles.button}>
          Send
        </button>
      </div>
    </div>
  );
}

export default ChatForm;