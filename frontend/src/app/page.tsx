import React from 'react';
import Navbar from './components/Navbar/Navbar';

export default function Home() {
  return (
    <main data-theme="dim">
      <Navbar />
      <div className="page-content">
        {/* Your page content goes here */}
        <button className="btn btn-ghost">Ghost</button>
      </div>
    </main>
  );
}