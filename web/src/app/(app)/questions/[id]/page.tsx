import React from 'react';
import { Answer, Question } from '@/types/models';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';
import AnswerForm from './answer-form';
import AddAnswerButton from './add-answer-button';
import Comment from './comment';

type Props = {
  params: Promise<{ id: number }>
};

const QuestionDetailPage: React.FC<Props> = async ({ params }) => {
  const { id } = await params
  const question: Question = await api.get(`questions/${id}`);
  const answers: Answer[] = await api.get(`questions/${id}/answers`)

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <span className={styles.levelBadge}>Lv {question.question_level}</span>
        <h1 className={styles.title}>{question.question_title}</h1>
      </div>

      <div className={styles.content}>
        <p>{question.question_content}</p>
      </div>
      {answers.length > 0 ? (
        <>
          {answers.map((answer, index) => (
            <div className={styles.answerSection} key={index}>
              <AnswerForm questionId={id} answer={answer} />
              <Comment answerId={answer.answer_id} />
            </div>
          ))}
          <AddAnswerButton questionId={id} />
        </>
      ) : (
        <div className={styles.answerSection}>
          <AnswerForm questionId={id} />
        </div>
      )}
    </div>
  );
};

export default QuestionDetailPage;