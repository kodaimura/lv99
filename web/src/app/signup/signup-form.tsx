'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { HttpError } from '@/lib/api/common';
import { api } from '@/lib/api/api.client';
import styles from './signup-form.module.css';

const SignupForm: React.FC = () => {
  const [name, setName] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [confirm_password, setConfirmPassword] = useState<string>('');
  const [error, setError] = useState<string>('');
  const router = useRouter();

  const handleSignup = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    if (password.length < 8) {
      setError('パスワードは8文字以上で入力してください。');
      return;
    } else if (password !== confirm_password) {
      setError('パスワードが一致しません。');
      return;
    }

    try {
      await api.post('accounts/signup', {
        name,
        password,
      });
      router.push('/login');
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
    <>
      <form className={styles.form} onSubmit={handleSignup}>
        {error && <p className={styles.error}>{error}</p>}
        <div className={styles.inputGroup}>
          <label htmlFor='name' className={styles.label}>
            アカウント名
          </label>
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
          <label htmlFor='password' className={styles.label}>
            パスワード
          </label>
          <input
            type='password'
            id='password'
            value={password}
            onChange={(e) => setPassword(e.target.value)}
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
    </>
  );
};

export default SignupForm;
