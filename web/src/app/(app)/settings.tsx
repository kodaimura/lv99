'use client';

import { useState, useEffect, useRef } from "react";
import Link from "next/link";
import { FiSettings } from "react-icons/fi";
import styles from "./settings.module.css";
import { api } from '@/lib/api/api.client';

const Settings = () => {
  const [open, setOpen] = useState(false);
  const containerRef = useRef<HTMLDivElement>(null);

  const logout = async () => {
    await api.post('/accounts/logout', {});
    window.location.replace('/login');
  }

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        containerRef.current &&
        !containerRef.current.contains(event.target as Node)
      ) {
        setOpen(false);
      }
    };

    if (open) {
      document.addEventListener("mousedown", handleClickOutside);
    } else {
      document.removeEventListener("mousedown", handleClickOutside);
    }

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [open]);

  return (
    <div className={styles.container} ref={containerRef}>
      <button
        className={styles.iconButton}
        onClick={() => setOpen(!open)}
      >
        <FiSettings className={styles.icon} />
      </button>

      {open && (
        <nav className={styles.menu}>
          <Link href="/account" className={styles.link}>アカウント設定</Link>
          <button className={styles.logout} onClick={logout}>ログアウト</button>
        </nav>
      )}
    </div>
  );
};

export default Settings;
