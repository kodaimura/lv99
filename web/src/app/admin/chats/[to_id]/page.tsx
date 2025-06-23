import styles from "./page.module.css";
import ChatArea from "./chat-area";

type Props = {
  params: { to_id: string }
};

const ChatPage: React.FC<Props> = async ({ params }) => {
  const { to_id } = params

  return (
    <ChatArea toId={parseInt(to_id)} />
  );
}

export default ChatPage;