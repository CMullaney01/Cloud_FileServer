import { NextResponse } from "next/server";
import { getAccessToken } from "../../utils/SessionTokenAccesor";
import { getServerSession } from "next-auth";
import { authOptions } from "../auth/[...nextauth]/route";

export async function POST(req) {
  const session = await getServerSession(authOptions);

  if (session) {
    const url = `${process.env.AUTH_BACKEND_URL}/api/v1/files/upload`; // Endpoint for file upload

    const postBody = await req.json();
    let accessToken = await getAccessToken();

    const resp = await fetch(url, {
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + accessToken,
        },
        method: "POST",
        body: JSON.stringify(postBody),
      });

    if (resp.ok) {
      const data = await resp.json();
      return NextResponse.json({ data }, { status: resp.status });
    }
    
    return NextResponse.json(
      { error: await resp.text() },
      { status: resp.status }
    );
  }
  return NextResponse.json({ error: "Unauthorized" }, { status: res.status });
}


export async function GET(req) {
  const session = await getServerSession(authOptions);

  if (session) {
    const next_url = new URL(req.url)
    const searchParams = new URLSearchParams(next_url.searchParams)
    const filename = searchParams.get('filename')
    const url = `${process.env.AUTH_BACKEND_URL}/api/v1/files/download?filename=${filename}`;

    let accessToken = await getAccessToken();

    const resp = await fetch(url, {
      headers: {
        Authorization: "Bearer " + accessToken,
      },
      method: "GET",
    });

    if (resp.ok) {
      const data = await resp.json();
      return NextResponse.json({ data }, { status: resp.status });
    }
    
    return NextResponse.json(
      { error: await resp.text() },
      { status: resp.status }
    );
  }
  
  return NextResponse.json({ error: "Unauthorized" }, { status: res.status });
}