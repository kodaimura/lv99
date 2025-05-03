import styles from "./page.module.css";
import Header from '@/app/components/layout/header';
import Main from '@/app/components/layout/main';
import Footer from "@/app/components/layout/footer";
import LoginLink from "@/app/components/parts/button/login-link";
import SignupLink from "@/app/components/parts/button/signup-link";

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
