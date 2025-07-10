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
              <Link href="/home" onClick={closeSidebar} className={styles.item}>ダッシュボード</Link>
            </li>
            <li>
              <Link href="/questions" onClick={closeSidebar} className={styles.item}>問題一覧</Link>
            </li>
            <li>
              <Link href="/documents" onClick={closeSidebar} className={styles.item}>ドキュメント</Link>
            </li>
            <li>
              <ChatMenu onClick={closeSidebar} />
            </li>
          </ul>
        </nav>
      </aside >
    </>
  );
};

export default Menu;
