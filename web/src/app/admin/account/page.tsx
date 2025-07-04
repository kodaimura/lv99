export const dynamic = 'force-dynamic';

import React from 'react';
import styles from './page.module.css';
import { api } from '@/lib/api/api.server';
import { Account, AccountProfile } from '@/types/models';
import AccountForm from './account-form';
import PasswordForm from './password-form';
import ProfileForm from './profile-form';
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: "lv99 - アカウント設定",
};

const AccountPage: React.FC = async () => {
  const account: Account = await api.get('accounts/me');
  const profile: AccountProfile = await api.get('accounts/me/profile');

  return (
    <div className={styles.container}>
      <AccountForm account={account} />
      <ProfileForm profile={profile} />
      <PasswordForm />
    </div>
  );
};

export default AccountPage;
