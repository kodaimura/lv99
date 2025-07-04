import type { Metadata } from "next";
import Header from "@/components/layout/header";
import Footer from "@/components/layout/footer";
import Main from "@/components/layout/main";
import Menu from "./menu";
import Settings from "./settings";

export const metadata: Metadata = {
  title: "lv99",
  description: "lv99は、問題を解きながらプログラミングスキルをレベルアップできる学習サービスです。",
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
      <Main style={{ paddingLeft: "240px" }}>
        {children}
        <Settings />
      </Main>
      <Footer />
    </>
  );
}
