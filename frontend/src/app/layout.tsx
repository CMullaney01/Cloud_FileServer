import type { Metadata } from "next";
import { Inter } from "next/font/google";
import './styles/globals.css';

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "File Server Frontend",
  description: "File server front end to interact with backend rest API",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" data-theme="dim">
      <body className={`p-4 ${inter.className}`}>{children}</body>
    </html>
  );
}
