import { Outlet } from 'react-router-dom';
import Header from '@/components/layout/header';
import Footer from '@/components/layout/footer';
import Main from '@/components/layout/main';
import Menu from './menu';
import styles from './layout.module.css';

export default function AdminLayout() {
  return (
    <>
      <Header><></></Header>
      <Menu />
      <Main className={styles.main}>
        <Outlet />
      </Main>
      <Footer />
    </>
  );
}
