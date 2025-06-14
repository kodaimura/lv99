"use client";

import React, { useState } from 'react';
import { HttpError } from '@/lib/api/common';
import { api } from '@/lib/api/api.client';
import styles from './password-form.module.css';
import { useRouter } from 'next/navigation';

const PasswordForm: React.FC = () => {
  const [currentPassword, setCurrentPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    if (newPassword.length < 8) {
      setError('パスワードは8文字以上で入力してください。');
      return;
    } else if (newPassword !== confirmPassword) {
      setError('パスワードが一致しません。');
      return;
    }

    try {
      await api.put('/accounts/me/password', {
        old_password: currentPassword,
        new_password: newPassword,
      });
      await api.post('/accounts/logout', {});
      router.push('/login');
    } catch (err) {
      if (err instanceof HttpError && err.status === 400) {
        setError('現在のパスワードが正しくありません。');
      } else {
        setError('パスワードの更新に失敗しました。\nもう一度お試しください。');
      }
    }
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      {error && <p className={styles.error}>{error}</p>}
      <div className={styles.inputGroup}>
        <label htmlFor="current-password" className={styles.label}>現在のパスワード</label>
        <input
          type="password"
          id="current-password"
          value={currentPassword}
          className={styles.input}
          onChange={(e) => setCurrentPassword(e.target.value)}
        />
      </div>
      <div className={styles.inputGroup}>
        <label htmlFor='new-password' className={styles.label}>新しいパスワード</label>
        <input
          type="password"
          id="new-password"
          value={newPassword}
          className={styles.input}
          onChange={(e) => setNewPassword(e.target.value)}
        />
      </div>
      <div className={styles.inputGroup}>
        <label htmlFor="confirm-password" className={styles.label}>確認用パスワード</label>
        <input
          type="password"
          id="confirm-password"
          value={confirmPassword}
          className={styles.input}
          onChange={(e) => setConfirmPassword(e.target.value)}
        />
      </div>
      <button type="submit" className={styles.submitButton}>更新</button>
    </form>
  );
}
export default PasswordForm;