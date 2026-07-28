import React, { useEffect, useState } from 'react';
import styles from './page.module.css';
import { api } from '@/lib/api/api.client';
import { Link } from 'react-router-dom';
import { CommentCount, Question, AnswerStatus } from '@/types/models';
import AnswerMatrix from './answer-matrix';
import LocalDate from '@/components/features/local-date';

const HomePage: React.FC = () => {
  const [questions, setQuestions] = useState<Question[]>([]);
  const [answerStatus, setAnswerStatus] = useState<AnswerStatus[]>([]);
  const [counts, setCounts] = useState<CommentCount[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;

    const fetchData = async () => {
      const since = new Date();
      since.setDate(since.getDate() - 7);

      try {
        const [questionsData, answerStatusData, countsData] = await Promise.all([
          api.get<Question[]>('questions'),
          api.get<AnswerStatus[]>('answers/status'),
          api.get<CommentCount[]>('comments/count', { since: since.toISOString().slice(0, 10) }),
        ]);

        if (!active) return;
        setQuestions(questionsData);
        setAnswerStatus(answerStatusData);
        setCounts(countsData);
      } catch (error) {
        console.error('Failed to fetch home data:', error);
      } finally {
        if (active) setLoading(false);
      }
    };

    fetchData();
    return () => {
      active = false;
    };
  }, []);

  if (loading) return <div className={styles.container}>読み込み中...</div>;

  return (
    <div className={styles.container}>
      <section className={styles.section}>
        <h2 className={styles.heading}>回答状況</h2>
        <AnswerMatrix answerStatus={answerStatus} questions={questions} />
      </section>

      <section className={styles.section}>
        <h2 className={styles.heading}>最近のコメント（7日以内）</h2>
        {counts.length > 0 && (
          <ul className={styles.countList}>
            {counts.map((count) => (
              <li key={count.answer_id} className={styles.countItem}>
                <Link to={`/questions/${count.question_level}`} className={styles.link}>
                  <div className={styles.countContent}>
                    <span className={styles.levelTag}>Lv {count.question_level}</span>
                    <span className={styles.title}>{count.question_title}</span>
                    <span className={styles.commentCount}>{count.comment_count} 件</span>
                    <span className={styles.date}>
                      <LocalDate isoString={count.created_at} />
                    </span>
                  </div>
                </Link>
              </li>
            ))}
          </ul>
        )}
      </section>
    </div>
  );
};

export default HomePage;
