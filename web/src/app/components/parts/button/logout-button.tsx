'use client';

import { api } from '@/app/lib/api/api.client';
import styles from './logout-button.module.css';

const logout = async () => {
  await api.post('/accounts/logout', {});
  window.location.replace('/login');
}

const LogoutButton = () => {
  return (
    <button className={styles.button} onClick={logout}>ログアウト</button>
  );
};

export default LogoutButton;