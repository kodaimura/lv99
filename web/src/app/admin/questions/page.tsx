'use client';

import React, { useEffect, useState } from 'react';
import { api } from '@/lib/api/api.client';
import styles from './page.module.css';
import Modal from '@/components/ui/modal';
import QuestionList from './question-list';
import QuestionForm from './question-form';
import type { Question } from "@/types/models";

const QuestionsPage: React.FC = () => {
  const [questions, setQuesions] = useState<Question[]>([]);
  const [question, setQuesion] = useState<Question | null>(null);
  const [showModal, setShowModal] = useState(false);

  useEffect(() => {
    getQuestions();
  }, [])

  const getQuestions = async () => {
    const data: Question[] = await api.get('admin/questions');
    setQuesions(data);
  }

  const handleSuccess = () => {
    getQuestions();
    setShowModal(false);
    setQuesion(null);
  };

  const handleClickEdit = (question: Question) => {
    setQuesion(question);
    setShowModal(true);
  }

  const handleClickDelete = async (question: Question) => {
    const questionId = question.id;
    try {
      await api.delete(`admin/questions/${questionId}`);
      getQuestions();
    } catch (err) {
      console.error('Error deleting question:', err);
      alert('問題の削除に失敗しました。');
    }
  }

  const handleClickRestore = async (question: Question) => {
    const questionId = question.id;
    try {
      await api.patch(`admin/questions/${questionId}`);
      getQuestions();
    } catch (err) {
      console.error('Error restoring question:', err);
      alert('問題の復元に失敗しました。');
    }
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <button onClick={() => setShowModal(true)} className={styles.addButton}>
          問題追加
        </button>
      </div>
      <Modal isOpen={showModal} onClose={() => setShowModal(false)} title='問題追加' >
        <QuestionForm onSuccess={handleSuccess} question={question} />
      </Modal>
      <QuestionList
        questions={questions}
        onClickEdit={handleClickEdit}
        onClickDelete={handleClickDelete}
        onClickRestore={handleClickRestore}
      />
    </div >
  );
};

export default QuestionsPage;
