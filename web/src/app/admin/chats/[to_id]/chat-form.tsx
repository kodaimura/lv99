'use client';

import { useEffect, useRef, useState } from "react";
import styles from "./chat-form.module.css";
import { Chat } from "@/types/models";

type Props = {
  toId: number;
  onRecieve: (chat: Chat) => void
};

const ChatForm: React.FC<Props> = ({ toId, onRecieve }) => {
  const [message, setMessage] = useState("");
  const socketRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8000/api/chats/ws");
    socketRef.current = socket;

    socket.onmessage = (event) => {
      onRecieve(JSON.parse(event.data));
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
  );
}

export default ChatForm;