import React from 'react';
import { Answer, Question, AccountWithProfile } from '@/types/models';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { oneDark } from 'react-syntax-highlighter/dist/esm/styles/prism';
import CommentList from './comment-list';
import LocalDate from '@/components/features/local-date';
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: "lv99 - 回答詳細",
};

type Props = {
  params: Promise<{ id: string }>;
};

const AnswerDetailPage: React.FC<Props> = async ({ params }) => {
  const { id } = await params;
  const answer: Answer = await api.get(`admin/answers/${id}`);
  const [account, question]: [AccountWithProfile, Question] = await Promise.all([
    api.get<AccountWithProfile>(`admin/accounts/${answer.account_id}/with-profile`),
    api.get<Question>(`questions/${answer.question_id}`),
  ]);

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
            {answer.is_correct ? "正解" : "不正解"}
          </span>
        </div>
        <div className={styles.metaRow}>
          <span className={styles.label}>正解日時:</span>
          <span>{answer.correct_at ? <LocalDate isoString={answer.correct_at} /> : "—"}</span>
        </div>
        <div className={styles.metaRow}>
          <span className={styles.label}>更新日時:</span>
          <span><LocalDate isoString={answer.updated_at} /></span>
        </div>

        <div className={styles.codeSection}>
          <h3 className={styles.codeTitle}>関数定義</h3>
          <SyntaxHighlighter language="python" style={oneDark} showLineNumbers>
            {answer.code_def || "# コードなし"}
          </SyntaxHighlighter>
        </div>

        <div className={styles.codeSection}>
          <h3 className={styles.codeTitle}>関数呼び出し</h3>
          <SyntaxHighlighter language="python" style={oneDark} showLineNumbers>
            {answer.code_call || "# コードなし"}
          </SyntaxHighlighter>
        </div>

        <div className={styles.codeSection}>
          <h3 className={styles.codeTitle}>出力</h3>
          <pre className={styles.callOutput}>
            {answer.call_output || "（出力なし）"}
          </pre>
        </div>

        <div className={styles.codeSection}>
          <h3 className={styles.codeTitle}>エラー</h3>
          <pre className={styles.callError}>
            {answer.call_error || "（エラーなし）"}
          </pre>
        </div>
      </div>

      <div className={styles.commentsSection}>
        <CommentList answerId={answer.id} />
      </div>
    </div>
  );
};

export default AnswerDetailPage;
