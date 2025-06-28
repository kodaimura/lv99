'use client';

import React, { useState } from 'react';
import styles from './comment-form.module.css';
import { api } from '@/lib/api/api.client';

type Props = {
  answerId: number;
  onSuccess?: () => void;
};

const CommentForm: React.FC<Props> = ({ answerId, onSuccess }) => {
  const [content, setContent] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    await api.post(`/comments`, {
      answer_id: answerId,
      content: content,
    });
    setContent('');
    setLoading(false);
    onSuccess?.();
  };

  return (
    <form onSubmit={handleSubmit} className={styles.form}>
      <textarea
        value={content}
        onChange={(e) => setContent(e.target.value)}
        required
        placeholder="コメントを入力..."
        className={styles.textarea}
        rows={3}
      />
      <button type="submit" disabled={loading} className={styles.button}>
        {loading ? '投稿中…' : 'コメントする'}
      </button>
    </form>
  );
};

export default CommentForm;
