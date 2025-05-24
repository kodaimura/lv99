import styles from "./page.module.css";
import ChatArea from "./chat-area";

type Props = {
  params: Promise<{ to_id: number }>
};

const ChatPage: React.FC<Props> = async ({ params }) => {
  const { to_id } = await params

  return (
    <ChatArea toId={to_id} />
  );
}

export default ChatPage;