import styles from "./page.module.css";
import Header from '@/components/layout/header';
import Main from '@/components/layout/main';
import Footer from "@/components/layout/footer";
import Link from 'next/link';

export default function Home() {
  return (
    <>
      <Header isLoggedIn={false}>
        <Link href="/login" className={styles.loginLink}>ログイン</Link>
      </Header>
      <Main>
        <div className={styles.container}>
          <section className={styles.intro}>
            <h1>lv99へようこそ</h1>
            <p>Pythonで問題を解き進めながら、初級〜中級のプログラミング力を鍛える学習サイトです。</p>
            <p>チャットで質問ができ、回答にはフィードバックも受けられます。</p>
            <p>興味のある方は
              <a href="mailto:contact@murakamikodai.com">contact@murakamikodai.com</a>
              までご連絡ください。
            </p>
            <p>軽い面談と説明を行い、問題の内容や進め方についてお話しします。</p>
            <p>話を聞きたいだけでも気軽にお問い合わせください。</p>
            <p>※お試し期間のため無料で学習できます。</p>
          </section>
        </div>
      </Main>
      <Footer />
    </>
  );
}
