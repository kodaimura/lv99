import React from 'react';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';
import QuestionList from './question-list';
import type { Question, AnswerStatus } from "@/types/models";

const QuestionsPage: React.FC = async () => {
  const questions: Question[] = await api.get('questions');
  const answerStatus: AnswerStatus[] = await api.get('answers/status');

  return (
    <div className={styles.container}>
      <QuestionList questions={questions} answerStatus={answerStatus} />
    </div >
  );
};

export default QuestionsPage;