export const dynamic = 'force-dynamic';

import React from 'react';
import { Metadata } from 'next';
import styles from './page.module.css';
import { api } from '@/lib/api/api.server';
import { Question, AnswerStatus } from '@/types/models';
import AnswerMatrix from './answer-matrix';

export const metadata: Metadata = {
  title: "lv99 - アカウント詳細",
};

type Props = {
  params: Promise<{ id: string }>;
};

const AccountDetailPage: React.FC<Props> = async ({ params }) => {
  const { id } = await params;

  const since = new Date();
  since.setDate(since.getDate() - 7);

  const questions: Question[] = await api.get('questions');
  const answerStatus: AnswerStatus[] = await api.get('/admin/answers/status', { account_id: id });

  return (
    <div className={styles.container}>
      <section className={styles.section}>
        <h2 className={styles.heading}>回答状況</h2>
        <AnswerMatrix
          accountId={parseInt(id)}
          answerStatus={answerStatus}
          questions={questions}
        />
      </section>
    </div>
  );
};

export default AccountDetailPage;
