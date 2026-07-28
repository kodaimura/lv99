import React, { useEffect, useState } from 'react';
import styles from './page.module.css';
import { api } from '@/lib/api/api.client';
import { Account, AccountProfile } from '@/types/models';
import AccountForm from './account-form';
import PasswordForm from './password-form';
import ProfileForm from './profile-form';

const AccountPage: React.FC = () => {
  const [account, setAccount] = useState<Account | null>(null);
  const [profile, setProfile] = useState<AccountProfile | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;

    const fetchData = async () => {
      try {
        const [accountData, profileData] = await Promise.all([
          api.get<Account>('accounts/me'),
          api.get<AccountProfile>('accounts/me/profile'),
        ]);
        if (!active) return;
        setAccount(accountData);
        setProfile(profileData);
      } catch (error) {
        console.error('Failed to fetch account settings:', error);
      } finally {
        if (active) setLoading(false);
      }
    };

    fetchData();
    return () => {
      active = false;
    };
  }, []);

  if (loading) return <div className={styles.container}>読み込み中...</div>;
  if (!account || !profile) return <div className={styles.container}>アカウント情報を取得できませんでした。</div>;

  return (
    <div className={styles.container}>
      <AccountForm account={account} />
      <ProfileForm profile={profile} />
      <PasswordForm />
    </div>
  );
};

export default AccountPage;
