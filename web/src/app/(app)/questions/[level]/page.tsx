import React from 'react';
import { Question } from '@/types/models';
import { api } from '@/lib/api/api.server';
import styles from './page.module.css';
import type { Metadata } from "next";
import AnswerList from './answer-list';
import Link from 'next/link';
import { IoIosArrowBack, IoIosArrowForward } from "react-icons/io";

export const metadata: Metadata = {
  title: "lv99 - 回答",
};

type Props = {
  params: Promise<{ level: string }>;
};

const QuestionDetailPage: React.FC<Props> = async ({ params }) => {
  const { level } = await params;
  let question: Question | null = null;
  try {
    // levelで取得に変更 ... levelはユニークである必要がある(運用で担保) 
    let questions: Question[] = await api.get(`questions?level=${level}`);
    question = questions[0];
  } catch { }

  if (!question) {
    return <div className={styles.container}>問題が見つかりません。</div>;
  }

  return (
    <div className={styles.container}>
      <div className={styles.navigation}>
        {question.level > 1 ?
          (
            <Link href={`/questions/${question.level - 1}`}>
              <IoIosArrowBack size={20} />
            </Link>
          ) : (
            <div style={{ opacity: 0.3 }}>
              <IoIosArrowBack size={20} />
            </div>
          )
        }
        <span className={styles.levelBadge}>Lv {question.level}</span>
        {question.id < 100 ?
          (
            <Link href={`/questions/${question.level + 1}`}>
              <IoIosArrowForward size={20} />
            </Link>
          ) : (
            <div style={{ opacity: 0.3 }}>
              <IoIosArrowForward size={20} />
            </div>
          )
        }
      </div>
      <div className={styles.header}>
        <h1 className={styles.title}>{question.title}</h1>
        <div className={styles.content}>
          <p>{question.content}</p>
        </div>
      </div>
      <AnswerList questionId={question.id} />
    </div>
  );
};

export default QuestionDetailPage;