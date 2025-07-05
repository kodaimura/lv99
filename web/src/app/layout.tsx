import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "@/globals.css";
import { Providers } from "./providers";
import type { Account } from "@/types/models";
import { cookies } from "next/headers";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "lv99",
  description: "lv99は、問題を段階的に解き進めるプログラミング学習サービスです。",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  async function getAccount(): Promise<Account | null> {
    try {
      const cookieStore = await cookies();
      const accessToken = cookieStore.get('access_token')?.value ?? "";

      const response = await fetch(`${process.env.API_HOST}/api/accounts/me`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Cookie: `access_token=${accessToken}`,
        },
        credentials: 'include',
      });

      if (!response.ok) {
        return null;
      }

      return (await response.json()) as Account;
    } catch {
      return null;
    }
  }

  let account: Account | null = null;
  try {
    account = await getAccount();
  } catch { }

  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable}`}>
        <Providers initAccount={account}>
          {children}
        </Providers>
      </body>
    </html>
  );
}
