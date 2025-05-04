'use client';

import React, { useEffect, useState } from 'react';
import { api } from '@/lib/api/api.client';
import styles from './page.module.css';
import Modal from '@/components/ui/Modal';
import QuestionList from './question-list';
import QuestionForm from './question-form';

const QuestionsPage: React.FC = () => {
  const [questions, setQuesions] = useState<any[]>([]);
  const [showModal, setShowModal] = useState(false);

  useEffect(() => {
    getQuestions();
  }, [])

  const getQuestions = async () => {
    const data: any[] = await api.get('questions');
    setQuesions(data);
  }

  const handleSuccess = () => {
    getQuestions();
    setShowModal(false);
  };

  return (
    <div className={styles.container}>
      <button onClick={() => setShowModal(true)} className={styles.addButton}>問題を追加</button>
      <Modal isOpen={showModal} onClose={() => setShowModal(false)} title='問題追加' >
        <QuestionForm onSuccess={handleSuccess} />
      </Modal>
      <QuestionList questions={questions} />
    </div >
  );
};

export default QuestionsPage;
