'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { HttpError } from '@/app/lib/api/common';
import { api } from '@/app/lib/api/api.client';
import styles from './page.module.css';

const LoginPage: React.FC = () => {
  const [admin_name, setAdminName] = useState('');
  const [admin_password, setAdminPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      await api.post('admins/login', { admin_name, admin_password });
      router.push('/admin/dashboard');
    } catch (err) {
      if (err instanceof HttpError && err.status === 401) {
        setError('ユーザ名又はパスワードが異なります。');
      } else {
        if (err instanceof Error) {
          console.error('ログインエラー:', err.message);
        }
        setError('ログインに失敗しました。');
      }
    }
  };

  return (
    <div className={styles.container}>
      <form className={styles.form} onSubmit={handleLogin}>
        {error && <p className={styles.error}>{error}</p>}
        <div className={styles.inputGroup}>
          <label htmlFor='admin_name' className={styles.label}>ユーザー名</label>
          <input
            type='text'
            id='admin_name'
            value={admin_name}
            onChange={(e) => setAdminName(e.target.value)}
            className={styles.input}
            required
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor='admin_password' className={styles.label}>パスワード</label>
          <input
            type='password'
            id='admin_password'
            value={admin_password}
            onChange={(e) => setAdminPassword(e.target.value)}
            className={styles.input}
            required
          />
        </div>
        <button type='submit' className={styles.submitButton}>ログイン</button>
      </form>
    </div>
  );
};

export default LoginPage;
