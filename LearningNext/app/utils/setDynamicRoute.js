'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';

// Might not be needed in later versions of next.js
// stops cache rendering in case a user becomes unauthenticated
export function SetDynamicRoute() {
  const router = useRouter();

  useEffect(() => {
    router.refresh();
  }, [router]);

  return <></>;
}