"use client";

import React, { useState } from 'react';
import { api } from '@/lib/api/api.client';
import styles from './password-form.module.css';
import { AccountProfile } from '@/types/models';

type Props = {
  profile: AccountProfile;
};

const ProfileForm: React.FC<Props> = ({ profile }) => {
  const [profileState, setProfileState] = useState<AccountProfile>(profile);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      await api.put('/accounts/me/profile', {
        display_name: profileState.display_name,
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
        <label htmlFor="display_name" className={styles.label}>表示名</label>
        <input
          type="text"
          id="display_name"
          value={profileState.display_name}
          className={styles.input}
          onChange={(e) => {
            setProfileState({ ...profileState, display_name: e.target.value });
          }}
        />
      </div>
      <button type="submit" className={styles.submitButton}>更新</button>
    </form>
  );
}
export default ProfileForm;