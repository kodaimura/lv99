import styles from "./page.module.css";
import Header from '@/components/layout/header';
import Main from '@/components/layout/main';
import Footer from "@/components/layout/footer";
import LoginLink from "@/components/parts/button/login-link";
import SignupLink from "@/components/parts/button/signup-link";

export default function Home() {
  return (
    <>
      <Header>
        <LoginLink /><SignupLink />
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
