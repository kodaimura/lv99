import React from 'react';
import styles from './page.module.css';
import Header from '@/components/layout/header';
import Main from '@/components/layout/main';
import Footer from '@/components/layout/footer';
import LoginForm from './login-form';
import { Link } from 'react-router-dom';

const LoginPage: React.FC = () => {
  return (
    <>
      <Header isLoggedIn={false}>
        <Link to="/" className={styles.homeLink}>トップへ戻る</Link>
      </Header>
      <Main>
        <div className={styles.container}>
          <LoginForm />
        </div>
      </Main>
      <Footer />
    </>
  );
};

export default LoginPage;
