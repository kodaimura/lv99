import React, { useEffect, useState } from 'react';
import { Answer, Question, AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.client';
import styles from './page.module.css';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { oneDark } from 'react-syntax-highlighter/dist/esm/styles/prism';
import CommentList from './comment-list';
import LocalDate from '@/components/features/local-date';
import { useParams } from 'react-router-dom';

const AnswerDetailPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [answer, setAnswer] = useState<Answer | null>(null);
  const [account, setAccount] = useState<AccountWithProfile | null>(null);
  const [question, setQuestion] = useState<Question | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;

    const fetchData = async () => {
      if (!id) {
        setLoading(false);
        return;
      }

      setLoading(true);
      try {
        const answerData: Answer = await api.get(`admin/answers/${id}`);
        const [accountData, questionData] = await Promise.all([
          api.get<AccountWithProfile>(`admin/accounts/${answerData.account_id}/with-profile`),
          api.get<Question>(`questions/${answerData.question_id}`),
        ]);

        if (!active) return;
        setAnswer(answerData);
        setAccount(accountData);
        setQuestion(questionData);
      } catch (error) {
        console.error('Failed to fetch answer detail:', error);
        if (!active) return;
        setAnswer(null);
        setAccount(null);
        setQuestion(null);
      } finally {
        if (active) setLoading(false);
      }
    };

    fetchData();
    return () => {
      active = false;
    };
  }, [id]);

  if (loading) return <div className={styles.container}>読み込み中...</div>;
  if (!answer || !account || !question) return <div className={styles.container}>回答情報を取得できませんでした。</div>;

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <span className={styles.levelBadge}>Lv {question.level}</span>
        <h1 className={styles.title}>{question.title}</h1>
        <div className={styles.content}>{question.content}</div>
      </div>

      <div className={styles.displayName}>回答者: {account.display_name}</div>

      <div className={styles.answerInfo}>
        <div className={styles.metaRow}>
          <span className={styles.label}>判定:</span>
          <span className={answer.is_correct ? styles.correct : styles.incorrect}>
            {answer.is_correct ? '正解' : '不正解'}
          </span>
        </div>
        <div className={styles.metaRow}>
          <span className={styles.label}>正解日時:</span>
          <span>{answer.correct_at ? <LocalDate isoString={answer.correct_at} /> : '—'}</span>
        </div>
        <div className={styles.metaRow}>
          <span className={styles.label}>更新日時:</span>
          <span><LocalDate isoString={answer.updated_at} /></span>
        </div>

        <div className={styles.codeSection}>
          <h3 className={styles.codeTitle}>関数定義</h3>
          <SyntaxHighlighter language="python" style={oneDark} showLineNumbers>
            {answer.code_def || '# コードなし'}
          </SyntaxHighlighter>
        </div>

        <div className={styles.codeSection}>
          <h3 className={styles.codeTitle}>関数呼び出し</h3>
          <SyntaxHighlighter language="python" style={oneDark} showLineNumbers>
            {answer.code_call || '# コードなし'}
          </SyntaxHighlighter>
        </div>

        <div className={styles.codeSection}>
          <h3 className={styles.codeTitle}>出力</h3>
          <pre className={styles.callOutput}>{answer.call_output || '（出力なし）'}</pre>
        </div>

        <div className={styles.codeSection}>
          <h3 className={styles.codeTitle}>エラー</h3>
          <pre className={styles.callError}>{answer.call_error || '（エラーなし）'}</pre>
        </div>
      </div>

      <div className={styles.commentsSection}>
        <CommentList answerId={answer.id} />
      </div>
    </div>
  );
};

export default AnswerDetailPage;
