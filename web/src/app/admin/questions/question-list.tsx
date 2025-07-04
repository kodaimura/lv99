import React from 'react';
import styles from './question-list.module.css';
import type { Question } from "@/types/models";
import LocalDate from '@/components/features/local-date';

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

  return (
    <div className={styles.tableContainer}>
      <table className={styles.table}>
        <thead className={styles.thead}>
          <tr>
            <th className={styles.th}>レベル</th>
            <th className={styles.th}>タイトル</th>
            <th className={styles.th}>内容</th>
            <th className={styles.th}>答え</th>
            <th className={styles.th}>更新日</th>
            <th className={styles.th}>更新</th>
            <th className={styles.th}>削除</th>
          </tr>
        </thead>
        <tbody className={styles.tbody}>
          {questions.map((q, i) => (
            <tr key={i} className={`${styles.tr} ${q.deleted_at ? styles.deletedRow : ''}`}>
              <td className={styles.td}>{q.level}</td>
              <td className={styles.td}>{q.title}</td>
              <td className={styles.td}>{q.content}</td>
              <td className={styles.td}>{q.answer}</td>
              <td className={styles.td}><LocalDate isoString={q.updated_at} /></td>
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
