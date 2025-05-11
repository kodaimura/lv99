import styles from "./page.module.css";
import ChatForm from "./chat-form";

type Props = {
  params: Promise<{ to_id: number }>
};

const ChatPage: React.FC<Props> = async ({ params }) => {
  const { to_id } = await params

  return (
    <ChatForm toId={to_id} />
  );
}

export default ChatPage;