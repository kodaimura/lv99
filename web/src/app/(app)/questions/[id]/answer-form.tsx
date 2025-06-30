'use client';

import React, { useState } from 'react';
import styles from './answer-form.module.css';
import { useRouter } from 'next/navigation';
import { api } from '@/lib/api/api.client';
import { Answer } from '@/types/models';
import { Trash2 } from 'lucide-react';
import { EditorView } from '@codemirror/view';
import CodeMirror from '@uiw/react-codemirror';
import { python } from '@codemirror/lang-python';
import { Play } from 'lucide-react';

const customTheme = EditorView.theme({
  '&': {
    fontSize: '19px',
    fontFamily: 'Fira Code, monospace',
    backgroundColor: '#f9fafb',
  },
});

type Props = {
  questionId: number;
  answer?: Answer;
};

const AnswerForm: React.FC<Props> = ({ questionId, answer }) => {
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const [id, setId] = useState<number | null>(answer?.id ?? null);
  const [codeDef, setCodeDef] = useState<string>(answer?.code_def ?? '');
  const [codeCall, setCodeCall] = useState<string>(answer?.code_call ?? '');
  const [isCorrect, setIsCorrect] = useState<boolean | null>(answer?.is_correct ?? null);
  const [callOutput, setCallOutput] = useState<string>(answer?.call_output ?? '');
  const [callError, setCallError] = useState<string>(answer?.call_error ?? '');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setIsCorrect(null);
    setCallOutput('');
    setCallError('');

    try {
      const response: Answer = id
        ? await api.put(`/answers/${id}`, { code_def: codeDef, code_call: codeCall })
        : await api.post(`/answers`, {
          question_id: questionId,
          code_def: codeDef,
          code_call: codeCall,
        });

      setId(response.id);
      setIsCorrect(response.is_correct);
      setCallOutput(response.call_output);
      setCallError(response.call_error);
    } catch (err) {
      console.error('Answer submission error:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async () => {
    if (!id || !confirm('この回答を削除しますか？')) return;
    await api.delete(`/answers/${id}`);
    router.refresh();
  };

  return (
    <form onSubmit={handleSubmit} className={styles.form}>
      {id && (
        <button
          type="button"
          onClick={handleDelete}
          disabled={loading}
          className={styles.deleteButton}
          aria-label="削除"
        >
          <Trash2 size={20} />
        </button>
      )}

      <h2 className={styles.heading}>関数定義</h2>
      <CodeMirror
        value={codeDef}
        extensions={[python(), customTheme]}
        onChange={(val) => setCodeDef(val)}
      />

      <h2 className={styles.heading}>関数呼出</h2>
      <CodeMirror
        value={codeCall}
        extensions={[python(), customTheme]}
        onChange={(val) => setCodeCall(val)}
      />

      <button type="submit" disabled={loading || !codeCall} className={styles.button}>
        <Play size={18} className={styles.icon} />
        {loading ? '実行中…' : '実行する'}
      </button>

      {(callOutput || callError) && (
        <div className={styles.outputSection}>
          {callOutput && (
            <>
              <h3 className={styles.subheading}>出力</h3>
              <pre className={styles.output}>{callOutput}</pre>
            </>
          )}
          {callError && (
            <>
              <h3 className={styles.subheading}>エラー</h3>
              <pre className={styles.error}>{callError}</pre>
            </>
          )}
        </div>
      )}

      {isCorrect !== null && (
        <p className={`${styles.result} ${isCorrect ? styles.correct : styles.incorrect}`}>
          {isCorrect ? '✅ 正解です！' : '❌ 不正解です'}
        </p>
      )}
    </form>
  );
};

export default AnswerForm;
