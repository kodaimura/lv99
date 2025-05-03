import styles from './login-link.module.css';
import Link from 'next/link';

const LoginLink = () => {
  return (
    <Link href="/login" className={styles.link}>ログイン</Link>
  );
};

export default LoginLink;