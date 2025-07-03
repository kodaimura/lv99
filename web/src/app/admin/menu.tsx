import React from 'react';
import styles from './menu.module.css';
import Link from 'next/link';
import ChatMenu from './chat-menu';

const Menu: React.FC = () => {
  return (
    <aside className={styles.sidebar}>
      <nav className={styles.menu}>
        <ul>
          <li>
            <Link href="/admin" className={styles.item}>ダッシュボード</Link>
          </li>
          <li>
            <Link href="/admin/questions" className={styles.item}>問題登録</Link>
          </li>
          <li>
            <Link href="/admin/answers" className={styles.item}>回答一覧</Link>
          </li>
          <li>
            <Link href="/admin/accounts" className={styles.item}>アカウント一覧</Link>
          </li>
        </ul>
      </nav>
      <hr />
      <ChatMenu />
    </aside >
  );
};

export default Menu;
