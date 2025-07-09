'use client';

import React, { useState, useEffect } from 'react';
import { Answer } from '@/types/models';
import { api } from '@/lib/api/api.client';
import styles from './answer-list.module.css';
import AnswerForm from './answer-form';
import CommentList from './comment-list';

type Props = {
  questionId: number;
};

const AnswerList: React.FC<Props> = ({ questionId }) => {
  const [answers, setAnswers] = useState<Answer[]>([]);
  const [showForm, setShowForm] = useState(false);

  const getAnswers = async () => {
    try {
      const response: Answer[] = await api.get("answers", { question_id: questionId });
      setAnswers(response);
      setShowForm(false);
    } catch (error) {
      console.error('Error fetching question or answers:', error);
      setAnswers([]);
    }
  };

  useEffect(() => {
    getAnswers();
  }, [questionId]);

  return (
    <>
      {
        answers.length > 0 ? (
          <>
            {answers.map((answer) => (
              <div className={styles.answerSection} key={answer.id}>
                <AnswerForm questionId={questionId} answer={answer} onDelete={() => getAnswers()} />
                <div className={styles.commentSection}>
                  <CommentList answerId={answer.id} />
                </div>
              </div>
            ))}
            <div className={styles.addAnswerButton}>
              {showForm && <AnswerForm questionId={questionId} onSubmit={() => getAnswers()} />}
            </div>
            <button onClick={() => setShowForm(true)} className={styles.button}>
              回答を追加する
            </button>
          </>
        ) : (
          <div className={styles.answerSection}>
            <AnswerForm questionId={questionId} onSubmit={() => getAnswers()} />
          </div>
        )
      }
    </>
  );
};

export default AnswerList;