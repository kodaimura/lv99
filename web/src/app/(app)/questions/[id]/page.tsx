import React from 'react';
import { Answer, Question } from '@/types/models';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';
import AnswerForm from './answer-form';
import AddAnswerButton from './add-answer-button';
import CommentList from './comment-list';
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "lv99 - 回答",
};

type Props = {
  params: Promise<{ id: string }>;
};

const QuestionDetailPage: React.FC<Props> = async ({ params }) => {
  const { id } = await params;
  const question: Question = await api.get(`questions/${id}`);
  const answers: Answer[] = await api.get("answers", { question_id: id });

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <span className={styles.levelBadge}>Lv {question.level}</span>
        <h1 className={styles.title}>{question.title}</h1>
        <div className={styles.content}>
          <p>{question.content}</p>
        </div>
      </div>
      {answers.length > 0 ? (
        <>
          {answers.map((answer, index) => (
            <div className={styles.answerSection} key={index}>
              <AnswerForm questionId={parseInt(id)} answer={answer} />
              <div className={styles.commentSection}>
                <CommentList answerId={answer.id} />
              </div>
            </div>
          ))}
          <AddAnswerButton questionId={parseInt(id)} />
        </>
      ) : (
        <div className={styles.answerSection}>
          <AnswerForm questionId={parseInt(id)} />
        </div>
      )}
    </div>
  );
};

export default QuestionDetailPage;