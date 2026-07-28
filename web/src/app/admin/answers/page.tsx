import React, { useEffect, useState } from 'react';
import styles from './page.module.css';
import { api } from '@/lib/api/api.client';
import AnswerList from './answer-list';
import { useSearchParams } from 'react-router-dom';

type AnswerSearchResult = {
  answer_id: number;
  updated_at: string;
  account_name: string;
  question_title: string;
  is_correct: boolean;
  code_def: string;
  comment_count: number;
  comment_at: string | null;
  comment_account_name: string | null;
};

const isIntegerString = (value: string | null): value is string => {
  if (typeof value !== 'string') return false;
  return /^-?\d+$/.test(value);
};

const AnswersPage: React.FC = () => {
  const [searchParams] = useSearchParams();
  const [answers, setAnswers] = useState<AnswerSearchResult[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;

    const fetchAnswers = async () => {
      setLoading(true);
      const params: Record<string, string> = {};
      const accountId = searchParams.get('account_id');
      const level = searchParams.get('level');

      if (isIntegerString(accountId)) params.account_id = accountId;
      if (isIntegerString(level)) params.level = level;

      try {
        const data = await api.get<AnswerSearchResult[]>('admin/answers/search', params);
        if (active) setAnswers(data);
      } catch (error) {
        console.error('Failed to fetch answers:', error);
      } finally {
        if (active) setLoading(false);
      }
    };

    fetchAnswers();
    return () => {
      active = false;
    };
  }, [searchParams]);

  if (loading) return <div className={styles.container}>読み込み中...</div>;

  return (
    <div className={styles.container}>
      <AnswerList answers={answers} />
    </div>
  );
};

export default AnswersPage;
