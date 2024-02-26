'use client';
import React from 'react'

import { SessionProvider } from 'next-auth/react';

// making the wrapper a client component
const SessionProviderWrapper = ({children}) => {
  return (
    <SessionProvider>{children}</SessionProvider>
  )
}

export default SessionProviderWrapper