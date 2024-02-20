"use client";

import { useSession, signIn, signOut } from "next-auth/react";
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
        Logged in as <span className="text-yellow-100">{session.user.email}</span>{" "}
        <button
          className="bg-blue-900 font-bold text-white py-1 px-2 rounded border border-gray-50"
          onClick={() => {
            keycloakSessionLogOut().then(() => signOut({ callbackUrl: "/" }));
          }}>
          Log out
        </button>
      </div>
    );
  }

  return (
    <div className="my-3">
      Not logged in.{" "}
      <button
        className="bg-blue-900 font-bold text-white py-1 px-2 rounded border border-gray-50"
        onClick={() => {
          console.log("AUTH_FRONTEND_CLIENT_ID:", process.env.AUTH_FRONTEND_CLIENT_ID);
          console.log("AUTH_FRONTEND_CLIENT_SECRET:", process.env.AUTH_FRONTEND_CLIENT_SECRET);
          console.log("AUTH_ISSUER:", process.env.AUTH_ISSUER);
          console.log("NEXTAUTH_URL:", process.env.NEXTAUTH_URL);
          console.log("NEXTAUTH_SECRET:", process.env.NEXTAUTH_SECRET);
          console.log("END_SESSION_URL:", process.env.END_SESSION_URL);
          console.log("AUTH_BACKEND_URL:", process.env.AUTH_BACKEND_URL);
          console.log("REFRESH_TOKEN_URL:", process.env.REFRESH_TOKEN_URL);
            // signIn("keycloak")
          }}>
        Log in
      </button>
    </div>
  );
}