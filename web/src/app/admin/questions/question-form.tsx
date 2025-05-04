'use client';

import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { HttpError } from '@/lib/api/common';
import { api } from '@/lib/api/api.client';
import styles from './question-form.module.css';
import type { Question } from "@/types/models";

type Props = {
  onSuccess: () => void;
  question?: Question;
};

const QuestionForm: React.FC<Props> = ({ onSuccess, question }) => {
  const [question_id, setQuestionId] = useState<number | null>(null);
  const [question_title, setQuestionTitle] = useState<string>('');
  const [question_content, setQuestionContent] = useState<string>('');
  const [question_answer, setQuestionAnswer] = useState<string>('');
  const [question_level, setQuestionLevel] = useState<number>(0);
  const [error, setError] = useState<string>('');
  const router = useRouter();

  useEffect(() => {
    if (question) {
      setQuestionId(question.question_id);
      setQuestionTitle(question.question_title);
      setQuestionContent(question.question_content);
      setQuestionAnswer(question.question_answer);
      setQuestionLevel(question.question_level);
    }
  }, [question]);

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    const body = {
      question_title,
      question_content,
      question_answer,
      question_level,
    }

    try {
      if (question_id) {
        await api.put(`admin/questions/${question_id}`, body);
      } else {
        await api.post('admin/questions', body);
      }
      onSuccess();
    } catch (err) {
      setError('登録に失敗しました。\nもう一度お試しください。');
    }
  };

  return (
    <>
      <form className={styles.form} onSubmit={submit}>
        {error && <p className={styles.error}>{error}</p>}
        <div className={styles.inputGroup}>
          <label htmlFor='question_title' className={styles.label}>
            タイトル
          </label>
          <input
            type='text'
            id='question_title'
            value={question_title}
            onChange={(e) => setQuestionTitle(e.target.value)}
            className={styles.input}
            required
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor='question_content' className={styles.label}>
            内容
          </label>
          <textarea
            id='question_content'
            value={question_content}
            onChange={(e) => setQuestionContent(e.target.value)}
            className={`${styles.input} ${styles.textarea}`}
            required
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor='question_answer' className={styles.label}>
            答え
          </label>
          <input
            type='text'
            id='question_answer'
            value={question_answer}
            onChange={(e) => setQuestionAnswer(e.target.value)}
            className={styles.input}
            required
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor='question_level' className={styles.label}>
            レベル
          </label>
          <input
            type='number'
            onChange={(e) => setQuestionLevel(parseInt(e.target.value))}
            value={question_level}
            className={styles.input}
            required
          />
        </div>
        <button type='submit' className={styles.submitButton}>
          登録
        </button>
      </form>
    </>
  );
};

export default QuestionForm;
