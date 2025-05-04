'use client';

import React from 'react';
import styles from './question-list.module.css';
import { api } from '@/lib/api/api.client';

type Props = {
  questions: any[];
  onClickRow: (q: any) => void
};

const QuestionList: React.FC<Props> = ({ questions, onClickRow }) => {
  const deleteRow = async (q: any) => {
    const question_id = q.question_id;
    try {
      await api.delete(`questions/${question_id}`);
    } catch (err) {

    }
  }

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
            <th className={styles.th}>更新</th>
            <th className={styles.th}>削除</th>
          </tr>
        </thead>
        <tbody className={styles.tbody}>
          {questions.map((q, i) => (
            <tr key={i} className={styles.tr}>
              <td className={styles.td}>{q.question_id}</td>
              <td className={styles.td}>{q.question_title}</td>
              <td className={styles.td}>{q.question_content}</td>
              <td className={styles.td}>{q.question_answer}</td>
              <td className={styles.td}>{q.question_level}</td>
              <td className={styles.td}>{q.updated_at}</td>
              <td className={styles.td}><button onClick={() => onClickRow(q)}>編集</button></td>
              <td className={styles.td}><button onClick={() => deleteRow(q)}>編集</button></td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default QuestionList;
