'use client';

import React, { useState } from 'react';
import styles from './answer-form.module.css';
import { api } from '@/lib/api/api.client';
import { Answer } from '@/types/models';

type Props = {
  questionId: number;
};

const AnswerForm: React.FC<Props> = ({ questionId }) => {
  const [code_def, setCodeDef] = useState('');
  const [code_call, setCodeCall] = useState('');
  const [is_correct, setIsCorrect] = useState<null | boolean>(null);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    const answer: Answer = await api.post(`/questions/${questionId}/answers`, {
      code_def, code_call
    });
    setIsCorrect(answer.is_correct);
    setLoading(false);
  };

  return (
    <form onSubmit={handleSubmit} className={styles.form}>
      <h2 className={styles.heading}>関数定義</h2>
      <textarea
        value={code_def}
        onChange={(e) => setCodeDef(e.target.value)}
        placeholder={`def add(a, b):\n    return a + b`}
        required
        className={styles.textarea}
        rows={15}
      />

      <h2 className={styles.heading}>関数呼出（答え）</h2>
      <input
        type="text"
        value={code_call}
        onChange={(e) => setCodeCall(e.target.value)}
        placeholder="add(1, 2)"
        required
        className={styles.input}
      />

      <button
        type="submit"
        disabled={loading}
        className={styles.button}
      >
        {loading ? '採点中…' : '解答する'}
      </button>

      {is_correct !== null && (
        <p className={`${styles.result} ${is_correct ? styles.correct : styles.incorrect}`}>
          {is_correct ? '✅ 正解です！' : '❌ 不正解です'}
        </p>
      )}
    </form>
  );
};

export default AnswerForm;
