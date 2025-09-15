'use client';

import React, { useCallback, useState, useEffect } from 'react';
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

  const getAnswers = useCallback(async () => {
    try {
      const response: Answer[] = await api.get("answers", { question_id: questionId });
      setAnswers(response);
      setShowForm(false);
    } catch (error) {
      console.error('Error fetching question or answers:', error);
      setAnswers([]);
    }
  }, [questionId]);

  useEffect(() => {
    getAnswers();
  }, [getAnswers]);

  return (
    <>
      {
        answers.length > 0 ? (
          <>
            <button onClick={() => setShowForm(true)} className={styles.button}>
              回答を追加する
            </button>
            <div>
              {showForm && <AnswerForm questionId={questionId} onSubmit={getAnswers} />}
            </div>
            {answers.map((answer, index) => (
              <div className={styles.answerSection} key={answer.id}>
                <AnswerForm questionId={questionId} answer={answer} onDelete={getAnswers} no={answers.length - index} />
                <div className={styles.commentSection}>
                  <CommentList answerId={answer.id} />
                </div>
              </div>
            ))}
          </>
        ) : (
          <div className={styles.answerSection}>
            <AnswerForm questionId={questionId} onSubmit={getAnswers} />
          </div>
        )
      }
    </>
  );
};

export default AnswerList;