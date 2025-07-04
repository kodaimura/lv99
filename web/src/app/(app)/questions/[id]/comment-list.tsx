'use client';

import React, { useEffect, useRef, useState } from 'react';
import { api } from '@/lib/api/api.client';
import styles from './comment-list.module.css';
import CommentForm from './comment-form';
import type { CommentWithProfile } from '@/types/models';
import LocalDate from '@/components/features/local-date';

type Props = {
  answerId: number;
};

const CommentList: React.FC<Props> = ({ answerId }) => {
  const [comments, setComments] = useState<CommentWithProfile[] | null>(null);
  const [isOpen, setIsOpen] = useState(false);
  const [openMenuIndex, setOpenMenuIndex] = useState<number | null>(null);
  const [editingCommentId, setEditingCommentId] = useState<number | null>(null);
  const menuRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    getComments();
  }, []);

  const getComments = async () => {
    const response: CommentWithProfile[] = await api.get("/comments/with-profile", { answer_id: answerId });
    setComments(response);
  };

  const handleToggleComments = () => {
    setIsOpen(prev => !prev);
    setEditingCommentId(null);
  };

  const handleMenuToggle = (index: number) => {
    setOpenMenuIndex(prev => (prev === index ? null : index));
    setEditingCommentId(null);
  };

  const handleStartEdit = (comment: CommentWithProfile) => {
    setEditingCommentId(comment.id);
    setOpenMenuIndex(null);
  };

  const handleCancelEdit = () => {
    setEditingCommentId(null);
  };

  const handleDelete = async (commentId: number) => {
    if (confirm('本当に削除しますか？')) {
      await api.delete(`/comments/${commentId}`);
      if (editingCommentId === commentId) {
        setEditingCommentId(null);
      }
      getComments();
    }
  };

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (menuRef.current && !menuRef.current.contains(e.target as Node)) {
        setOpenMenuIndex(null);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  return (
    <div>
      <button onClick={handleToggleComments} className={styles.commentButton}>
        {`${comments?.length ?? 0} 件のコメント`}
      </button>

      {isOpen && comments && (
        <>
          <div className={styles.commentList}>
            {comments.map((comment, index) => (
              <div key={comment.id} className={styles.commentItem}>
                <div className={styles.commentHeader}>
                  <div className={styles.commentMeta}>
                    <span className={styles.commentDate}>
                      <LocalDate isoString={comment.created_at} />
                    </span>
                    <span className={styles.commentAuthor}>{comment.display_name}</span>
                  </div>

                  <button
                    className={styles.menuButton}
                    onClick={() => handleMenuToggle(index)}
                    aria-haspopup="true"
                    aria-expanded={openMenuIndex === index}
                  >
                    ⋮
                  </button>

                  {openMenuIndex === index && (
                    <div className={styles.menu} ref={menuRef}>
                      <div className={styles.menuItem} onClick={() => handleStartEdit(comment)}>更新</div>
                      <div className={styles.menuItem} onClick={() => handleDelete(comment.id)}>削除</div>
                    </div>
                  )}
                </div>

                {editingCommentId === comment.id ? (
                  <CommentForm
                    answerId={answerId}
                    commentId={comment.id}
                    initialContent={comment.content}
                    onSuccess={() => {
                      setEditingCommentId(null);
                      getComments();
                    }}
                    onCancel={handleCancelEdit}
                  />
                ) : (
                  <p className={styles.commentContent}>{comment.content}</p>
                )}
              </div>
            ))}
          </div>

          {editingCommentId === null && (
            <CommentForm answerId={answerId} onSuccess={getComments} />
          )}
        </>
      )}
    </div>
  );
};

export default CommentList;
