import styles from "./page.module.css";
import Header from '@/components/layout/header';
import Main from '@/components/layout/main';
import Footer from "@/components/layout/footer";
import Link from 'next/link';

export default function Home() {
  return (
    <>
      <Header>
        <Link href="/login" className={styles.loginLink}>ログイン</Link>
        <Link href="/signup" className={styles.signupLink}>サインアップ</Link>
      </Header>
      <Main>
        <div className={styles.container}>
          <section className={styles.intro}>
            <h2>Getting Started</h2>
            <p>Next.js makes building web applications easy and efficient.</p>
          </section>
        </div>
      </Main>
      <Footer />
    </>
  );
}
