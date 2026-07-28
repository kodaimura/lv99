'use client';

import React from 'react';
import styles from './answer-matrix.module.css';
import type { AnswerStatus, Question } from '@/types/models';
import { Link } from 'react-router-dom';

type Props = {
  questions: Question[];
  answerStatus: AnswerStatus[];
};

const AnswerMatrix: React.FC<Props> = ({ questions, answerStatus }) => {
  const levelMap: Record<number, number> = {};
  const questionToLevelMap: Record<number, number> = {};

  questions.forEach((q) => {
    questionToLevelMap[q.id] = q.level;
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
        const hasQuestion = questions.some((question) => question.level === level);
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

        return hasQuestion ? (
          <Link key={level} to={`/questions/${level}`} className={styles.link}>
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
