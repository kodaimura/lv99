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
  params: Promise<{ id: string }>;
};

const QuestionDetailPage: React.FC<Props> = async ({ params }) => {
  const { id } = await params;
  let question: Question | null = null;
  try {
    question = await api.get(`questions/${id}`);
  } catch { }

  if (!question) {
    return <div className={styles.container}>問題が見つかりません。</div>;
  }

  return (
    <div className={styles.container}>
      <div className={styles.navigation}>
        {question.id > 1 ?
          (
            <Link href={`/questions/${question.id - 1}`}>
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
            <Link href={`/questions/${question.id + 1}`}>
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