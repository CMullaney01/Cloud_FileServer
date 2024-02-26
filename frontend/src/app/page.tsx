"use client";

import React from 'react';
import Navbar from './components/Navbar/Navbar';
import ArchitectureCard from './components/ArchitectureCard/ArchitectureCard'
import DashCard from './components/DashCard/DashCard'
import TutorialCard from './components/TutorialCard/TutorialCard'
export default function Home() {
  return (
    <main>
      <Navbar />
      <div className="page-content mt-12 grid grid-cols-1 md:grid-cols-3 gap-8">
        {/* Architecture Cards */}
        <ArchitectureCard />
        <DashCard />
        <TutorialCard  />
      </div>
    </main>
  );
}