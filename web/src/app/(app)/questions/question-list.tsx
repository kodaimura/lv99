'use client'

import React, { useState, useEffect } from 'react';
import styles from './question-list.module.css';
import { useRouter } from 'next/navigation';
import type { Question, AnswerStatus } from "@/types/models";
import LocalDate from '@/components/features/local-date';
import { api } from '@/lib/api/api.client';

const QuestionList: React.FC = () => {
  const [questions, setQuestions] = useState<Question[]>([]);
  const [statusMap, setStatusMap] = useState<Record<number, AnswerStatus>>([]);

  const router = useRouter();
  const getQuestions = async () => {
    try {
      const response: Question[] = await api.get('questions');
      setQuestions(response);
    } catch (error) {
      console.error("Failed to fetch questions:", error);
    }
  };

  const getAnswerStatus = async () => {
    try {
      const answerStatus: AnswerStatus[] = await api.get('answers/status');
      answerStatus.forEach((status) => {
        setStatusMap((prev) => ({
          ...prev,
          [status.question_id]: status,
        }));
      });
    } catch (error) {
      console.error("Failed to fetch questions:", error);
    }
  };

  useEffect(() => {
    getQuestions();
    getAnswerStatus();
  }, []);

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
                {statusMap[q.id] && (statusMap[q.id].correct_at ? (<LocalDate isoString={statusMap[q.id].correct_at} />) : "-")}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div >
  );
};

export default QuestionList;
