'use client';

import React, { useState } from 'react';
import AnswerForm from './answer-form';
import styles from './add-answer-button.module.css';

type Props = {
  questionId: number;
};

const AddAnswerButton: React.FC<Props> = ({ questionId }) => {
  const [showForm, setShowForm] = useState(false);

  return (
    <div className={styles.addAnswerButton}>
      {showForm ? (
        <AnswerForm questionId={questionId} />
      ) : (
        <button onClick={() => setShowForm(true)} className={styles.button}>
          回答を追加する
        </button>
      )}
    </div>
  );
};

export default AddAnswerButton;