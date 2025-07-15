import styles from './header.module.css';

interface Props {
  isLoggedIn?: boolean;
  children: React.ReactNode;
}

const Header: React.FC<Props> = ({ isLoggedIn = true, children }) => {
  return (
    <header className={`${styles.header} ${isLoggedIn ? styles.loggedIn : ''}`}>
      <h1 className={styles.title}>lv99</h1>
      <div className={styles.childrenWrapper}>
        {children}
      </div>
    </header>
  );
};

export default Header;