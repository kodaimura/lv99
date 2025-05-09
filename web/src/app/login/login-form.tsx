'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { HttpError } from '@/lib/api/common';
import { api } from '@/lib/api/api.client';
import styles from './login-form.module.css';

const LoginForm: React.FC = () => {
  const [name, setName] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [error, setError] = useState<string>('');
  const router = useRouter();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    try {
      const response: any = await api.post('accounts/login', { name, password });
      if (response?.account?.role == 0) {
        router.push('/admin');
      } else {
        router.push('/home');
      }
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
    <>
      <form className={styles.form} onSubmit={handleLogin}>
        {error && <p className={styles.error}>{error}</p>}
        <div className={styles.inputGroup}>
          <label htmlFor='name' className={styles.label}>ユーザー名</label>
          <input
            type='text'
            id='name'
            value={name}
            onChange={(e) => setName(e.target.value)}
            className={styles.input}
            required
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor='password' className={styles.label}>パスワード</label>
          <input
            type='password'
            id='password'
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className={styles.input}
            required
          />
        </div>
        <button type='submit' className={styles.submitButton}>ログイン</button>
      </form>
    </>
  );
};

export default LoginForm;
