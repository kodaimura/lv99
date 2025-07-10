'use client';

import React, { useState } from 'react';
import styles from './menu.module.css';
import Link from 'next/link';
import ChatMenu from './chat-menu';

const Menu: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggleSidebar = () => setIsOpen(!isOpen);
  const closeSidebar = () => setIsOpen(false);

  return (
    <>
      <button className={styles.menuButton} onClick={toggleSidebar}>
        ☰
      </button>

      {isOpen && <div className={styles.overlay} onClick={closeSidebar}></div>}

      <aside className={`${styles.sidebar} ${isOpen ? styles.open : ''}`}>
        <nav className={styles.menu}>
          <ul>
            <li>
              <Link href="/admin" onClick={closeSidebar} className={styles.item}>ダッシュボード</Link>
            </li>
            <li>
              <Link href="/admin/questions" onClick={closeSidebar} className={styles.item}>問題登録</Link>
            </li>
            <li>
              <Link href="/admin/answers" onClick={closeSidebar} className={styles.item}>回答一覧</Link>
            </li>
            <li>
              <Link href="/admin/accounts" onClick={closeSidebar} className={styles.item}>アカウント一覧</Link>
            </li>
          </ul>
        </nav>
        <hr />
        <ChatMenu onClick={closeSidebar} />
      </aside>
    </>
  );
};

export default Menu;
