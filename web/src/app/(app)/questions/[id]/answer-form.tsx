'use client';

import React, { useState } from 'react';
import styles from './answer-form.module.css';
import { api } from '@/lib/api/api.client';
import { Answer } from '@/types/models';

type Props = {
  questionId: number;
  answer?: Answer;
};

const AnswerForm: React.FC<Props> = ({ questionId, answer }) => {
  const answerId = answer?.answer_id;
  const [code_def, setCodeDef] = useState<string>(answer?.code_def ?? '');
  const [code_call, setCodeCall] = useState<string>(answer?.code_call ?? '');
  const [is_correct, setIsCorrect] = useState<null | boolean>(answer?.is_correct ?? null);
  const [correct_at, setCorrectAt] = useState<null | string>(answer?.correct_at ?? null);
  const [call_output, setCallOutput] = useState<string>(answer?.call_output ?? '');
  const [call_error, setCallError] = useState<string>(answer?.call_error ?? '');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    let response: Answer;
    if (answerId) {
      response = await api.put(`/questions/${questionId}/answers/${answerId}`, {
        code_def, code_call
      });
    } else {
      response = await api.post(`/questions/${questionId}/answers`, {
        code_def, code_call
      });
    }
    setIsCorrect(response.is_correct);
    setCallOutput(response.call_output);
    setCallError(response.call_error);

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

      <h2 className={styles.heading}>関数呼出</h2>
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

      {call_output && (
        <div className={styles.outputSection}>
          <h3 className={styles.subheading}>出力</h3>
          <pre className={styles.output}>{call_output}</pre>
        </div>
      )}

      {call_error && (
        <div className={styles.errorSection}>
          <h3 className={styles.subheading}>エラー</h3>
          <pre className={styles.error}>{call_error}</pre>
        </div>
      )}

      {is_correct !== null && (
        <p className={`${styles.result} ${is_correct ? styles.correct : styles.incorrect}`}>
          {is_correct ? '✅ 正解です！' : '❌ 不正解です'}
        </p>
      )}
    </form>
  );
};

export default AnswerForm;
