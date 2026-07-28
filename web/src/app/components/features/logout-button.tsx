'use client';

import { api } from '@/lib/api/api.client';
import { clearAccessToken } from '@/lib/auth-token';
import styles from './logout-button.module.css';

const logout = async () => {
  try {
    await api.post('/accounts/logout', {});
  } finally {
    clearAccessToken();
    window.location.replace('/login');
  }
}

const LogoutButton = () => {
  return (
    <button className={styles.button} onClick={logout}>ログアウト</button>
  );
};

export default LogoutButton;
