import React from 'react';
import styles from './question-list.module.css';
import type { Question } from "@/types/models";

type Props = {
  questions: Question[];
  onClickEdit: (question: Question) => void
  onClickDelete: (question: Question) => void
  onClickRestore: (question: Question) => void
};

const QuestionList: React.FC<Props> = ({
  questions,
  onClickEdit,
  onClickDelete,
  onClickRestore
}) => {

  const formatDate = (dateStr: string) => {
    const date = new Date(dateStr);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}`;
  };

  return (
    <div className={styles.tableContainer}>
      <table className={styles.table}>
        <thead className={styles.thead}>
          <tr>
            <th className={styles.th}>#</th>
            <th className={styles.th}>タイトル</th>
            <th className={styles.th}>内容</th>
            <th className={styles.th}>答え</th>
            <th className={styles.th}>レベル</th>
            <th className={styles.th}>更新日</th>
            <th className={styles.th}>更新</th>
            <th className={styles.th}>削除</th>
          </tr>
        </thead>
        <tbody className={styles.tbody}>
          {questions.map((q, i) => (
            <tr key={i} className={styles.tr}>
              <td className={styles.td}>{q.question_id}</td>
              <td className={styles.td}>{q.question_title}</td>
              <td className={styles.td}>{q.question_content}</td>
              <td className={styles.td}>{q.question_answer}</td>
              <td className={styles.td}>{q.question_level}</td>
              <td className={styles.td}>{formatDate(q.updated_at)}</td>
              <td className={styles.td}><button onClick={() => onClickEdit(q)}>編集</button></td>
              <td className={styles.td}>
                {
                  q.deleted_at ?
                    <button onClick={() => onClickRestore(q)}>復元</button> :
                    <button onClick={() => onClickDelete(q)}>削除</button>
                }
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default QuestionList;
