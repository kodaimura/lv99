import React from 'react';
import styles from './menu.module.css';
import Link from 'next/link';

const Menu: React.FC = () => {
  return (
    <aside className={styles.sidebar}>
      <nav className={styles.menu}>
        <ul>
          <li><Link href="/home" className={styles.link}>ダッシュボード</Link></li>
          <li><Link href="/questions" className={styles.link}>問題一覧</Link></li>
          <li><Link href="/chat" className={styles.link}>チャット</Link></li>
          <li><Link href="/documents" className={styles.link}>ドキュメント</Link></li>
        </ul>
      </nav>
    </aside >
  );
};

export default Menu;
