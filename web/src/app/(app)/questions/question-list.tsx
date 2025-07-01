'use client'

import React from 'react';
import styles from './question-list.module.css';
import type { AnswerStatus, Question } from "@/types/models";
import { useRouter } from 'next/navigation';

type Props = {
  questions: Question[];
  answerStatus: AnswerStatus[];
};

const QuestionList: React.FC<Props> = ({ questions, answerStatus }) => {
  const router = useRouter();
  const statusMap: Record<number, AnswerStatus> = {};
  answerStatus.forEach((status) => {
    statusMap[status.question_id] = status;
  });

  const formatDate = (dateStr: string | null) => {
    if (!dateStr) return null;
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
            <th className={styles.th}>Lv.</th>
            <th className={styles.th}>タイトル</th>
            <th className={styles.th}></th>
            <th className={styles.th}>正解数</th>
            <th className={styles.th}>正解日時</th>
          </tr>
        </thead>
        <tbody className={styles.tbody}>
          {questions.map((q, i) => (
            <tr key={i} className={styles.tr} onClick={() => router.push(`/questions/${q.id}`)}>
              <td className={styles.td}>{q.level}</td>
              <td className={styles.td}>{q.title}</td>
              <td className={styles.td}>
                {statusMap[q.id] && (statusMap[q.id].is_correct ? "✅" : "❌")}
              </td>
              <td className={styles.td}>
                {statusMap[q.id] && (statusMap[q.id].correct_count > 0 ? statusMap[q.id].correct_count : "-")}
              </td>
              <td className={styles.td}>
                {statusMap[q.id] && (statusMap[q.id].correct_at ? formatDate(statusMap[q.id].correct_at) : "-")}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div >
  );
};

export default QuestionList;
