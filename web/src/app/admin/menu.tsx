import React from 'react';
import styles from './menu.module.css';
import Link from 'next/link';

const Menu: React.FC = () => {
  return (
    <aside className={styles.sidebar}>
      <nav className={styles.menu}>
        <ul>
          <li><Link href="/admin" className={styles.link}>ダッシュボード</Link></li>
          <li><Link href="/admin/questions" className={styles.link}>問題一覧</Link></li>
          <li><Link href="/admin/chats" className={styles.link}>チャット</Link></li>
          <li><Link href="/admin/answers" className={styles.link}>回答一覧</Link></li>
        </ul>
      </nav>
    </aside >
  );
};

export default Menu;
