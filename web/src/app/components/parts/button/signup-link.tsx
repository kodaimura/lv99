import styles from './login-link.module.css';
import Link from 'next/link';

const SignupLink = () => {
  return (
    <Link href="/signup" className={styles.link}>サインアップ</Link>
  );
};

export default SignupLink;