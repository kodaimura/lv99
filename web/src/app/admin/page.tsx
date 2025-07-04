export const dynamic = 'force-dynamic';

import React from 'react';
import styles from './page.module.css';
import { api } from '@/lib/api/api.server';
import Link from 'next/link';
import { CommentCount } from '@/types/models';
import LocalDate from '@/components/features/local-date';
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: "lv99 - HOME",
};

const AdminPage: React.FC = async () => {
  const since = new Date();
  since.setDate(since.getDate() - 7);

  const counts: CommentCount[] = await api.get('admin/comments/count', { since: since.toISOString().slice(0, 10) });

  return (
    <div className={styles.container}>
      {counts.length > 0 && (
        <ul className={styles.countList}>
          {counts.map((count) => (
            <li key={count.answer_id} className={styles.countItem}>
              <Link href={`admin/answers/${count.answer_id}`} className={styles.link}>
                <div className={styles.countContent}>
                  <span className={styles.levelTag}>Lv {count.question_level}</span>
                  <span className={styles.title}>{count.question_title}</span>
                  <span className={styles.commentCount}>{count.comment_count} 件のコメント</span>
                  <span className={styles.date}>
                    <LocalDate isoString={count.created_at} />
                  </span>
                </div>
              </Link>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default AdminPage;
