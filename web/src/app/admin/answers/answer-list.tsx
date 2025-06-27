'use client';

import React from 'react';
import styles from './answer-list.module.css';
import { useRouter } from 'next/navigation';

type Props = {
  answers: any[];
};

const AnswerList: React.FC<Props> = ({
  answers,
}) => {
  const router = useRouter();

  const formatDate = (dateStr: string) => {
    const date = new Date(dateStr);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}`;
  };

  return (
    <div className={styles.tableContainer}>
      <table className={styles.table}>
        <thead className={styles.thead}>
          <tr>
            <th className={styles.th}>#</th>
            <th className={styles.th}>回答者名</th>
            <th className={styles.th}>問題</th>
            <th className={styles.th}>正誤</th>
            <th className={styles.th}>更新日時</th>
            <th className={styles.th}>コメント</th>
            <th className={styles.th}>コメント日時</th>
            <th className={styles.th}>コメント者名</th>
          </tr>
        </thead>
        <tbody className={styles.tbody}>
          {answers.map((a, i) => (
            <tr key={i} className={styles.tr} onClick={() => { router.push(`answers/${a.answer_id}`); }}>
              <td className={styles.td}>{a.answer_id}</td>
              <td className={styles.td}>{a.account_name}</td>
              <td className={styles.td}>{a.question_title}</td>
              <td className={styles.td}>
                <div className={styles.tooltip}>
                  {a.is_correct ? '🟢' : '❌'}
                  <div className={styles.tooltipText}>{a.code_def}</div>
                </div>
              </td>
              <td className={styles.td}>{formatDate(a.updated_at)}</td>
              <td className={styles.td}>{a.comment_count}</td>
              <td className={styles.td}>{a.comment_at ? formatDate(a.comment_at) : '-'}</td>
              <td className={styles.td}>{a.comment_account_name || '-'}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default AnswerList;
