import type { Metadata } from "next";
import { Inter } from "next/font/google";
import './styles/globals.css';
import Footer from './components/Footer/Footer'
import SessionProviderWrapper from '@/app/utils/SessionProviderWrapper'

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
    <SessionProviderWrapper>
      <html lang="en" data-theme="dim">
        <body className={inter.className}>
            <div className="p-4 flex-grow">
              {children}
            </div>
        </body>
        <Footer />
      </html>
    </SessionProviderWrapper>
  );
}
