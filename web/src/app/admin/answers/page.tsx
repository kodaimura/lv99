import React from 'react';
import styles from './page.module.css';
import { api } from '@/lib/api/api.server';
import AnswerList from './answer-list';
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: "lv99 - 回答一覧",
};
type SearchParams = {
  account_id?: string;
  level?: string;
};

const AnswersPage = async ({
  searchParams
}: {
  searchParams: Promise<SearchParams>;
}) => {
  const { account_id, level } = await searchParams;
  const isIntegerString = (value: string | any): boolean => {
    if (typeof value !== 'string') {
      return false;
    }
    return /^-?\d+$/.test(value);
  }
  let params = {};
  if (isIntegerString(account_id)) {
    params = { ...params, account_id: account_id };
  }
  if (isIntegerString(level)) {
    params = { ...params, level: level };
  }

  const data: any[] = await api.get('admin/answers/search', params);
  return (
    <div className={styles.container}>
      <AnswerList answers={data} />
    </div>
  );
};

export default AnswersPage;
