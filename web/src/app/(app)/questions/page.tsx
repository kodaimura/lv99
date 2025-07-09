export const dynamic = 'force-dynamic';

import React from 'react';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';
import QuestionList from './question-list';
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "lv99 - 問題一覧",
};

const QuestionsPage: React.FC = async () => {
  return (
    <div className={styles.container}>
      <QuestionList />
    </div >
  );
};

export default QuestionsPage;