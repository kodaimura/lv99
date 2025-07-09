import React from 'react';
import { Question } from '@/types/models';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';
import type { Metadata } from "next";
import AnswerList from './answer-list';

export const metadata: Metadata = {
  title: "lv99 - 回答",
};

type Props = {
  params: Promise<{ id: string }>;
};

const QuestionDetailPage: React.FC<Props> = async ({ params }) => {
  const { id } = await params;
  const question: Question = await api.get(`questions/${id}`);

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <span className={styles.levelBadge}>Lv {question.level}</span>
        <h1 className={styles.title}>{question.title}</h1>
        <div className={styles.content}>
          <p>{question.content}</p>
        </div>
      </div>
      <AnswerList questionId={parseInt(id)} />
    </div>
  );
};

export default QuestionDetailPage;