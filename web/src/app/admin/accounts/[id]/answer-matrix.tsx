'use client';

import React from 'react';
import styles from './answer-matrix.module.css';
import type { AnswerStatus, Question } from '@/types/models';
import Link from 'next/link';

type Props = {
  accountId: number;
  questions: Question[];
  answerStatus: AnswerStatus[];
};

const AnswerMatrix: React.FC<Props> = ({ accountId, questions, answerStatus }) => {
  const levelMap: Record<number, number> = {};
  const levelToQuestionIdMap: Record<number, number> = {};
  const questionToLevelMap: Record<number, number> = {};

  questions.forEach((q) => {
    questionToLevelMap[q.id] = q.level;
    levelToQuestionIdMap[q.level] = q.id;
  });

  answerStatus.forEach((status) => {
    const level = questionToLevelMap[status.question_id];
    if (level != null) {
      levelMap[level] = (levelMap[level] || 0) + status.correct_count;
    }
  });

  const levels = Array.from({ length: 100 }, (_, i) => i + 1);

  const getColor = (count: number) => {
    if (count === 0) return '#ffffff';
    if (count === 1) return '#fff7cc';
    if (count === 2) return '#ffee99';
    if (count === 3) return '#ffe066';
    if (count === 4) return '#ffd633';
    return '#ffcc00';
  };

  return (
    <div className={styles.grid}>
      {levels.map((level) => {
        const count = levelMap[level] || 0;
        const questionId = levelToQuestionIdMap[level];
        const cell = (
          <div
            key={level}
            className={styles.cell}
            style={{ backgroundColor: getColor(count) }}
            title={`Lv ${level} (${count})`}
          >
            {level}
          </div>
        );

        return questionId ? (
          <Link key={level} href={`/admin/answers/?account_id=${accountId}&level=${level}`} className={styles.link}>
            {cell}
          </Link>
        ) : (
          cell
        );
      })}
    </div>
  );
};

export default AnswerMatrix;
