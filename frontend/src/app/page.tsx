import React from 'react';
import Navbar from './components/Navbar/Navbar';
import architectureImage from '../../public/file-storage.png'; // Import your PNG file
import Image from 'next/image';

export default function Home() {
  return (
    <main data-theme="dim">
      <Navbar />
      <div className="page-content">
        {/* Your page content goes here */}
        <Image src={architectureImage} alt="Your Image" /> {/* Use the imported image */}
        <button className="btn btn-ghost">Ghost</button>
      </div>
    </main>
  );
}