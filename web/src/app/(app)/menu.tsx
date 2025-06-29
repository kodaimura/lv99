import React from 'react';
import styles from './menu.module.css';
import Link from 'next/link';
import UnreadCountMap from './unread-count';

const Menu: React.FC = () => {
  return (
    <aside className={styles.sidebar}>
      <nav className={styles.menu}>
        <ul>
          <li>
            <Link href="/home" className={styles.item}>ダッシュボード</Link>
          </li>
          <li>
            <Link href="/questions" className={styles.item}>問題一覧</Link>
          </li>
          <li>
            <Link href="/chat" className={styles.item}>
              チャット<UnreadCountMap />
            </Link>
          </li>
          <li>
            <Link href="/account" className={styles.item}>アカウント設定</Link>
          </li>
          <li>
            <Link href="/documents" className={styles.item}>ドキュメント</Link>
          </li>
        </ul>
      </nav>
    </aside >
  );
};

export default Menu;
