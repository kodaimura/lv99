'use client';

import React from 'react';
import styles from './form.module.css';
import { Account } from '@/types/models';

type Props = {
  account: Account;
};

const AccountForm: React.FC<Props> = ({ account }) => {
  return (
    <form className={styles.form} onSubmit={(e) => e.preventDefault()}>
      <div className={styles.inputGroup}>
        <label className={styles.label}>ログインユーザ名</label>
        <input
          type="text"
          disabled
          value={account.name}
          className={styles.input}
        />
      </div>
    </form>
  );
}
export default AccountForm;