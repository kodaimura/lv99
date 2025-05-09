'use client';

import React, { useEffect, useState } from 'react';
import { api } from '@/lib/api/api.client';
import styles from './comment.module.css';
import CommentForm from './comment-form';
import type { Comment } from '@/types/models';

type Props = {
  answerId: number;
};

const CommentCount: React.FC<Props> = ({ answerId }) => {
  const [comments, setComments] = useState<Comment[] | null>(null);
  const [isOpen, setIsOpen] = useState(false);

  useEffect(() => {
    getComments();
  }, [])

  const getComments = async () => {
    const response: Comment[] = await api.get(`answers/${answerId}/comments`);
    setComments(response);
  }

  const handleToggleComments = async () => {
    setIsOpen((prev) => !prev);
  };

  return (
    <div>
      <button onClick={handleToggleComments} className={styles.commentButton}>
        {`${comments?.length ?? 0} 件のコメント`}
      </button>

      {isOpen && comments && (
        <>
          <div className={styles.commentList}>
            {comments.map((comment, index) => (
              <div key={index} className={styles.commentItem}>
                <p>{comment.comment_content}</p>
              </div>
            ))}
          </div>
          <CommentForm answerId={answerId} onSuccess={getComments} />
        </>
      )}
    </div>
  );
};

export default CommentCount;
