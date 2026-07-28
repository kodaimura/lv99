import React from 'react';
import styles from './page.module.css';
import QuestionList from './question-list';

const QuestionsPage: React.FC = () => {
  return (
    <div className={styles.container}>
      <QuestionList />
    </div>
  );
};

export default QuestionsPage;
