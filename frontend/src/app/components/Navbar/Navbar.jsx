import React from 'react';
import Link from 'next/link';
import AuthStatus from '@/app/components/AuthStatus/AuthStatus'

const Navbar = () => {
  return (
      <div className="navbar bg-base-300 rounded-lg shadow-md p-4 text-neutral-content">
        <div className="navbar-start">
          <div className="dropdown">
            <div tabIndex={0} role="button" className="btn btn-ghost">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8m-8 6h16" /></svg>
            </div>
            <ul tabIndex="0" class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
              <li><Link Link href="/Dashboard" >Dashboard</Link></li>
              <li>
                <a>About</a>
                <ul class="p-2">
                  <li><Link href="/" >Tutorial</Link></li>
                  <li><Link href="/" >Docs</Link></li>
                </ul>
              </li>
            </ul>
          </div>
          <Link className="btn btn-ghost text-xl" href="/">
            Conor Mullaney&apos;s File Server
          </Link>
        </div>
        <div className="navbar-end">
          <AuthStatus />
        </div>
      </div>
  );
};

export default Navbar;