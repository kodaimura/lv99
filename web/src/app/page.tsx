import styles from "./page.module.css";

export default function Home() {
  return (
    <div className={styles.container}>
      <section className={styles.intro}>
        <h2>Getting Started</h2>
        <p>Next.js makes building web applications easy and efficient.</p>
      </section>
    </div>
  );
}
