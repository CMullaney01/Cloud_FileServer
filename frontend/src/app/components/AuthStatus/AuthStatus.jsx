"use client";

import { useSession, signIn, signOut } from "next-auth/react";
import Link from "next/link";
import { useEffect } from "react";

async function keycloakSessionLogOut() {
    try {
      await fetch(`/api/auth/logout`, { method: "GET" });
    } catch (err) {
      console.error(err);
    }
  }

export default function AuthStatus() {
  const { data: session, status } = useSession(); 

  useEffect(() => {
    
    if (
      status != "loading" &&
      session &&
      session?.error === "RefreshAccessTokenError"
    ) {
      signOut({ callbackUrl: "/" });
    }
  }, [session, status]);


  if (status == "loading") {
    return <div className="my-3">Loading...</div>;
  } else if (session) {
    return (
      <div className="my-3">
        Hello <span>{session.user.name}!</span>{" "}
        <div
          className="btn"
          onClick={() => {
            // end keycloak
            keycloakSessionLogOut().then(() => signOut({ callbackUrl: "/" }));
          }}>
          Log out
        </div>
      </div>
    );
  }

  return (
    <div>
      <div className="btn mr-2"  onClick={() => {
        signIn("keycloak")
      }}>
        Login/Sign-up
      </div>
    </div>
  );
}
