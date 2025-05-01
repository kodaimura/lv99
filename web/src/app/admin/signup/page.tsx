'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { HttpError } from '@/app/lib/api/common';
import { api } from '@/app/lib/api/api.client';
import styles from './page.module.css';

const SignupPage: React.FC = () => {
  const [admin_name, setAdminName] = useState('');
  const [admin_password, setAdminPassword] = useState('');
  const [confirm_password, setConfirmPassword] = useState('');
  const [error, setError] = useState('');
  const router = useRouter();

  const handleSignup = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    if (admin_password.length < 8) {
      setError('パスワードは8文字以上で入力してください。');
      return;
    } else if (admin_password !== confirm_password) {
      setError('パスワードが一致しません。');
      return;
    }

    try {
      await api.post('admins/signup', {
        admin_name,
        admin_password,
      });
      router.push('/admin/login');
    } catch (err) {
      if (err instanceof HttpError && err.status === 409) {
        setError('ユーザ名が既に使われています。');
      } else {
        if (err instanceof Error) {
          console.error('サインアップ失敗', err);
        }
        setError('サインアップに失敗しました。\nもう一度お試しください。');
      }
    }
  };

  return (
    <div className={styles.container}>
      <form className={styles.form} onSubmit={handleSignup}>
        {error && <p className={styles.error}>{error}</p>}
        <div className={styles.inputGroup}>
          <label htmlFor='admin_name' className={styles.label}>
            アカウント名
          </label>
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
          <label htmlFor='admin_password' className={styles.label}>
            パスワード
          </label>
          <input
            type='password'
            id='admin_password'
            value={admin_password}
            onChange={(e) => setAdminPassword(e.target.value)}
            className={styles.input}
            required
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor='confirm_password' className={styles.label}>
            パスワード（確認）
          </label>
          <input
            type='password'
            id='confirm_password'
            value={confirm_password}
            onChange={(e) => setConfirmPassword(e.target.value)}
            className={styles.input}
            required
          />
        </div>
        <button type='submit' className={styles.submitButton}>
          サインアップ
        </button>
      </form>
    </div>
  );
};

export default SignupPage;
