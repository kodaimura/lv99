import React, { useEffect, useState } from 'react';
import { Question } from '@/types/models';
import { api } from '@/lib/api/api.client';
import styles from './page.module.css';
import AnswerList from './answer-list';
import { Link, useParams } from 'react-router-dom';
import { IoIosArrowBack, IoIosArrowForward } from 'react-icons/io';

const QuestionDetailPage: React.FC = () => {
  const { level } = useParams<{ level: string }>();
  const [question, setQuestion] = useState<Question | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;

    const fetchQuestion = async () => {
      setLoading(true);
      try {
        const questions: Question[] = await api.get('questions', { level });
        if (active) setQuestion(questions[0] ?? null);
      } catch {
        if (active) setQuestion(null);
      } finally {
        if (active) setLoading(false);
      }
    };

    fetchQuestion();
    return () => {
      active = false;
    };
  }, [level]);

  if (loading) return <div className={styles.container}>読み込み中...</div>;
  if (!question) return <div className={styles.container}>問題が見つかりません。</div>;

  return (
    <div className={styles.container}>
      <div className={styles.navigation}>
        {question.level > 1 ? (
          <Link to={`/questions/${question.level - 1}`}>
            <IoIosArrowBack size={20} />
          </Link>
        ) : (
          <div style={{ opacity: 0.3 }}>
            <IoIosArrowBack size={20} />
          </div>
        )}
        <span className={styles.levelBadge}>Lv {question.level}</span>
        {question.id < 100 ? (
          <Link to={`/questions/${question.level + 1}`}>
            <IoIosArrowForward size={20} />
          </Link>
        ) : (
          <div style={{ opacity: 0.3 }}>
            <IoIosArrowForward size={20} />
          </div>
        )}
      </div>
      <div className={styles.header}>
        <h1 className={styles.title}>{question.title}</h1>
        <div className={styles.content}>
          <p>{question.content}</p>
        </div>
      </div>
      <AnswerList questionId={question.id} />
    </div>
  );
};

export default QuestionDetailPage;
