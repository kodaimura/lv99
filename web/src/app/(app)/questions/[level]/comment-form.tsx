'use client';

import React, { useState, useEffect } from 'react';
import styles from './comment-form.module.css';
import { api } from '@/lib/api/api.client';

type Props = {
  answerId: number;
  commentId?: number;
  initialContent?: string;
  onSuccess?: () => void;
  onCancel?: () => void;
};

const CommentForm: React.FC<Props> = ({
  answerId,
  commentId,
  initialContent = '',
  onSuccess,
  onCancel,
}) => {
  const [content, setContent] = useState(initialContent);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setContent(initialContent);
  }, [initialContent]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!content.trim()) {
      alert('コメントを入力してください。');
      return;
    }
    setLoading(true);

    if (commentId) {
      await api.put(`/comments/${commentId}`, { content });
    } else {
      await api.post(`/comments`, {
        answer_id: answerId,
        content,
      });
    }

    setContent('');
    setLoading(false);
    onSuccess?.();
  };

  return (
    <form onSubmit={handleSubmit} className={styles.form}>
      <textarea
        value={content}
        onChange={(e) => setContent(e.target.value)}
        placeholder="コメントを入力..."
        className={styles.textarea}
        rows={7}
      />
      <div className={styles.buttonGroup}>
        <button type="submit" disabled={loading} className={styles.button}>
          {loading ? (commentId ? '更新中…' : '投稿中…') : (commentId ? '更新する' : 'コメントする')}
        </button>
        {commentId && onCancel && (
          <button
            type="button"
            onClick={onCancel}
            className={styles.buttonCancel}
            disabled={loading}
          >
            キャンセル
          </button>
        )}
      </div>
    </form>
  );
};

export default CommentForm;
