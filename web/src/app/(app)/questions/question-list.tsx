import React from 'react';
import styles from './question-list.module.css';
import type { Question } from "@/types/models";
import Link from 'next/link';

type Props = {
  questions: Question[];
};

const QuestionList: React.FC<Props> = ({ questions }) => {

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
            <th className={styles.th}>タイトル</th>
            <th className={styles.th}>内容</th>
            <th className={styles.th}>答え</th>
            <th className={styles.th}>レベル</th>
            <th className={styles.th}>更新日</th>
          </tr>
        </thead>
        <tbody className={styles.tbody}>
          {questions.map((q, i) => (
            <tr key={i} className={styles.tr}>
              <td className={styles.td}>{q.id}</td>
              <td className={styles.td}><Link href={`questions/${q.id}`}>{q.title}</Link></td>
              <td className={styles.td}>{q.content}</td>
              <td className={styles.td}>{q.answer}</td>
              <td className={styles.td}>{q.level}</td>
              <td className={styles.td}>{formatDate(q.updated_at)}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div >
  );
};

export default QuestionList;
