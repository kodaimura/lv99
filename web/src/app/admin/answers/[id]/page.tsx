import React from 'react';
import { AccountWithProfile, Answer, Question } from '@/types/models';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';

type Props = {
  params: Promise<{ id: string }>
};

const AnswerDetailPage: React.FC<Props> = async ({ params }) => {
  const { id } = await params;
  const account: AccountWithProfile = await api.get(`admin/accounts/${id}/with-profile`);
  const answers: Answer[] = await api.get("admin/answers", { account_id: id });
  const questions: Question[] = await api.get(`admin/questions`);

  console.log(account);
  console.log(answers);
  console.log(questions);
  return (
    <div className={styles.container}>
      <div className={styles.header}>
      </div>

      <div className={styles.content}>
      </div>
    </div>
  );
};

export default AnswerDetailPage;