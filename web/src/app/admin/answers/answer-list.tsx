'use client';

import React from 'react';
import styles from './answer-list.module.css';
import { useRouter } from 'next/navigation';
import LocalDate from '@/components/features/local-date';

type Props = {
  answers: any[];
};

const AnswerList: React.FC<Props> = ({
  answers,
}) => {
  const router = useRouter();

  return (
    <div className={styles.tableContainer}>
      <table className={styles.table}>
        <thead className={styles.thead}>
          <tr>
            <th className={styles.th}>更新日時</th>
            <th className={styles.th}>回答者名</th>
            <th className={styles.th}>問題</th>
            <th className={styles.th}>正誤</th>
            <th className={styles.th}>コメント</th>
            <th className={styles.th}>コメント日時</th>
            <th className={styles.th}>コメント者名</th>
          </tr>
        </thead>
        <tbody className={styles.tbody}>
          {answers.map((a, i) => (
            <tr key={i} className={styles.tr} onClick={() => { router.push(`answers/${a.answer_id}`); }}>
              <td className={styles.td}><LocalDate isoString={a.updated_at} /></td>
              <td className={styles.td}>{a.account_name}</td>
              <td className={styles.td}>{a.question_title}</td>
              <td className={styles.td}>
                <div className={styles.tooltip}>
                  {a.is_correct ? '🟢' : '❌'}
                  <div className={styles.tooltipText}>{a.code_def}</div>
                </div>
              </td>
              <td className={styles.td}>{a.comment_count}</td>
              <td className={styles.td}>{a.comment_at ? <LocalDate isoString={a.comment_at} /> : '-'}</td>
              <td className={styles.td}>{a.comment_account_name || '-'}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default AnswerList;
