import type { Metadata } from "next";
import Header from "@/components/layout/header";
import Footer from "@/components/layout/footer";
import Main from "@/components/layout/main";
import Menu from "./menu";
import styles from "./layout.module.css";

export const metadata: Metadata = {
  title: "lv99",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <Header><></></Header>
      <Menu />
      <Main className={styles.main}>
        {children}
      </Main>
      <Footer />
    </>
  );
}
