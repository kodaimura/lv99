'use client';

import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { api } from '@/lib/api/api.client';
import styles from './question-form.module.css';
import type { Question } from "@/types/models";

type Props = {
  onSuccess: () => void;
  question: Question | null;
};

const QuestionForm: React.FC<Props> = ({ onSuccess, question }) => {
  const [id, setId] = useState<number | null>(null);
  const [title, setTitle] = useState<string>('');
  const [content, setContent] = useState<string>('');
  const [answer, setAnswer] = useState<string>('');
  const [level, setLevel] = useState<number>(0);
  const [error, setError] = useState<string>('');
  const router = useRouter();

  useEffect(() => {
    if (question) {
      setId(question.id);
      setTitle(question.title);
      setContent(question.content);
      setAnswer(question.answer);
      setLevel(question.level);
    }
  }, [question]);

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    const body = {
      title,
      content,
      answer,
      level,
    }

    try {
      if (id) {
        await api.put(`admin/questions/${id}`, body);
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
          <label htmlFor='title' className={styles.label}>
            タイトル
          </label>
          <input
            type='text'
            id='title'
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className={styles.input}
            required
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor='content' className={styles.label}>
            内容
          </label>
          <textarea
            id='content'
            value={content}
            onChange={(e) => setContent(e.target.value)}
            className={`${styles.input} ${styles.textarea}`}
            required
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor='answer' className={styles.label}>
            答え
          </label>
          <input
            type='text'
            id='answer'
            value={answer}
            onChange={(e) => setAnswer(e.target.value)}
            className={styles.input}
            required
          />
        </div>
        <div className={styles.inputGroup}>
          <label htmlFor='level' className={styles.label}>
            レベル
          </label>
          <input
            type='number'
            onChange={(e) => setLevel(parseInt(e.target.value))}
            value={level}
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
