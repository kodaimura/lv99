"use client";

import React, { useState } from 'react';
import { HttpError } from '@/lib/api/common';
import { api } from '@/lib/api/api.client';
import styles from './password-form.module.css';
import { useRouter } from 'next/navigation';
import { AccountProfile } from '@/types/models';

type Props = {
  accountProfile: AccountProfile;
};

const ProfileForm: React.FC<Props> = ({ accountProfile }) => {
  const [profile, setProfile] = useState<AccountProfile>(accountProfile);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      await api.put('/accounts/me/profile', {
        display_name: profile.display_name,
      });
      alert('プロフィールを更新しました。');
    } catch (err) {
      setError('更新に失敗しました。\nもう一度お試しください。');
    }
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      {error && <p className={styles.error}>{error}</p>}
      <div className={styles.inputGroup}>
        <label htmlFor="display_name">表示名</label>
        <input
          type="text"
          id="display_name"
          value={profile.display_name}
          className={styles.input}
          onChange={(e) => {
            setProfile({ ...profile, display_name: e.target.value });
          }}
        />
      </div>
      <button type="submit" className={styles.submitButton}>更新</button>
    </form>
  );
}
export default ProfileForm;