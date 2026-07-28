import React, { useEffect, useState } from 'react';
import styles from './page.module.css';
import { api } from '@/lib/api/api.client';
import { Link } from 'react-router-dom';
import { CommentCount } from '@/types/models';
import LocalDate from '@/components/features/local-date';

const AdminPage: React.FC = () => {
  const [counts, setCounts] = useState<CommentCount[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let active = true;

    const fetchCounts = async () => {
      const since = new Date();
      since.setDate(since.getDate() - 7);

      try {
        const data: CommentCount[] = await api.get('admin/comments/count', {
          since: since.toISOString().slice(0, 10),
        });
        if (active) setCounts(data);
      } catch (error) {
        console.error('Failed to fetch admin dashboard:', error);
      } finally {
        if (active) setLoading(false);
      }
    };

    fetchCounts();
    return () => {
      active = false;
    };
  }, []);

  if (loading) return <div className={styles.container}>読み込み中...</div>;

  return (
    <div className={styles.container}>
      {counts.length > 0 && (
        <ul className={styles.countList}>
          {counts.map((count) => (
            <li key={count.answer_id} className={styles.countItem}>
              <Link to={`/admin/answers/${count.answer_id}`} className={styles.link}>
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
    </div>
  );
};

export default AdminPage;
