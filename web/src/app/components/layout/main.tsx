import styles from './main.module.css';

interface Props {
  style?: React.CSSProperties;
  className?: string;
  children: React.ReactNode;
}

const Main: React.FC<Props> = ({ style, className, children }) => {
  return (
    <main className={`${styles.main} ${className}`} style={style}>
      {children}
    </main>
  );
};

export default Main;