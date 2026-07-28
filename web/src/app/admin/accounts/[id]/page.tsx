import React, { useEffect, useMemo, useState } from 'react';
import styles from './page.module.css';
import { api } from '@/lib/api/api.client';
import { Question, AnswerStatus } from '@/types/models';
import AnswerMatrix from './answer-matrix';
import { useParams } from 'react-router-dom';

const AccountDetailPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const accountId = useMemo(() => Number.parseInt(id ?? '', 10), [id]);
  const [questions, setQuestions] = useState<Question[]>([]);
  const [answerStatus, setAnswerStatus] = useState<AnswerStatus[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;

    const fetchData = async () => {
      if (!Number.isInteger(accountId)) {
        setLoading(false);
        return;
      }

      setLoading(true);
      try {
        const [questionsData, answerStatusData] = await Promise.all([
          api.get<Question[]>('questions'),
          api.get<AnswerStatus[]>('/admin/answers/status', { account_id: accountId }),
        ]);

        if (!active) return;
        setQuestions(questionsData);
        setAnswerStatus(answerStatusData);
      } catch (error) {
        console.error('Failed to fetch account detail:', error);
      } finally {
        if (active) setLoading(false);
      }
    };

    fetchData();
    return () => {
      active = false;
    };
  }, [accountId]);

  if (loading) return <div className={styles.container}>読み込み中...</div>;
  if (!Number.isInteger(accountId)) return <div className={styles.container}>アカウントが見つかりません。</div>;

  return (
    <div className={styles.container}>
      <section className={styles.section}>
        <h2 className={styles.heading}>回答状況</h2>
        <AnswerMatrix
          accountId={accountId}
          answerStatus={answerStatus}
          questions={questions}
        />
      </section>
    </div>
  );
};

export default AccountDetailPage;
