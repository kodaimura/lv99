export const dynamic = 'force-dynamic';

import React from 'react';
import { Metadata } from 'next';
import styles from './page.module.css';
import { api } from '@/lib/api/api.server';
import Link from 'next/link';
import { CommentCount, Question, AnswerStatus } from '@/types/models';
import AnswerMatrix from './answer-matrix';
import LocalDate from '@/components/features/local-date';

export const metadata: Metadata = {
  title: "lv99 - HOME",
};

const HomePage: React.FC = async () => {
  const since = new Date();
  since.setDate(since.getDate() - 7);

  const questions: Question[] = await api.get('questions');
  const answerStatus: AnswerStatus[] = await api.get('answers/status');
  const counts: CommentCount[] = await api.get('comments/count', { since: since.toISOString().slice(0, 10) });

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
                <Link href={`questions/${count.question_id}`} className={styles.link}>
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
