import React from 'react';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';
import QuestionList from './question-list';
import type { Question } from "@/types/models";

const QuestionsPage: React.FC = async () => {
  const questions: Question[] = await api.get('questions');

  return (
    <div className={styles.container}>
      <QuestionList questions={questions} />
    </div >
  );
};

export default QuestionsPage;