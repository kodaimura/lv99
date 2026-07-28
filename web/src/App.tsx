import { Navigate, Outlet, Route, Routes, useLocation } from 'react-router-dom';
import { useAuth } from '@/contexts/auth-context';
import PublicHome from '@/page';
import LoginPage from '@/login/page';
import AppLayout from '@/(app)/layout';
import HomePage from '@/(app)/home/page';
import QuestionsPage from '@/(app)/questions/page';
import QuestionDetailPage from '@/(app)/questions/[level]/page';
import DocumentsPage from '@/(app)/documents/page';
import ChatPage from '@/(app)/chat/page';
import AccountPage from '@/(app)/account/page';
import AdminLayout from '@/admin/layout';
import AdminHomePage from '@/admin/page';
import AdminQuestionsPage from '@/admin/questions/page';
import AdminAnswersPage from '@/admin/answers/page';
import AdminAnswerDetailPage from '@/admin/answers/[id]/page';
import AdminAccountsPage from '@/admin/accounts/page';
import AdminAccountDetailPage from '@/admin/accounts/[id]/page';
import AdminChatPage from '@/admin/chats/[to_id]/page';
import AdminAccountPage from '@/admin/account/page';
import { getAuthenticatedEntryPath } from '@/lib/auth-redirect';

const LoadingScreen = () => <div style={{ padding: '4rem 1.5rem', textAlign: 'center' }}>読み込み中...</div>;

const PublicOnly = () => {
  const { account, isLoading } = useAuth();

  if (isLoading) return <LoadingScreen />;
  if (account) return <Navigate to={getAuthenticatedEntryPath(account)} replace />;

  return <Outlet />;
};

const RequireAuth = () => {
  const { account, isLoading } = useAuth();
  const location = useLocation();

  if (isLoading) return <LoadingScreen />;
  if (!account) return <Navigate to="/login" replace state={{ from: location }} />;

  return <Outlet />;
};

const RequireAdmin = () => {
  const { account, isLoading } = useAuth();
  const location = useLocation();

  if (isLoading) return <LoadingScreen />;
  if (!account) return <Navigate to="/login" replace state={{ from: location }} />;
  if (account.role !== 0) return <Navigate to="/home" replace />;

  return <Outlet />;
};

export default function App() {
  return (
    <Routes>
      <Route element={<PublicOnly />}>
        <Route path="/" element={<PublicHome />} />
        <Route path="/login" element={<LoginPage />} />
      </Route>

      <Route element={<RequireAuth />}>
        <Route element={<AppLayout />}>
          <Route path="/home" element={<HomePage />} />
          <Route path="/questions" element={<QuestionsPage />} />
          <Route path="/questions/:level" element={<QuestionDetailPage />} />
          <Route path="/documents" element={<DocumentsPage />} />
          <Route path="/chat" element={<ChatPage />} />
          <Route path="/account" element={<AccountPage />} />
        </Route>
      </Route>

      <Route element={<RequireAdmin />}>
        <Route element={<AdminLayout />}>
          <Route path="/admin" element={<AdminHomePage />} />
          <Route path="/admin/questions" element={<AdminQuestionsPage />} />
          <Route path="/admin/answers" element={<AdminAnswersPage />} />
          <Route path="/admin/answers/:id" element={<AdminAnswerDetailPage />} />
          <Route path="/admin/accounts" element={<AdminAccountsPage />} />
          <Route path="/admin/accounts/:id" element={<AdminAccountDetailPage />} />
          <Route path="/admin/chats/:to_id" element={<AdminChatPage />} />
          <Route path="/admin/account" element={<AdminAccountPage />} />
        </Route>
      </Route>

      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  );
}
