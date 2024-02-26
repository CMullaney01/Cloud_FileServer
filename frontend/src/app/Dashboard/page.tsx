import { getServerSession } from "next-auth";
import { authOptions } from "../api/auth/[...nextauth]/route";
import { redirect } from "next/navigation";
import { getAccessToken } from "@/app/utils/SessionTokenAccesor";
import { SetDynamicRoute } from "@/app/utils/SetDynamicRoute";
import FileTree from "@/app/components/FileTree/FileTree"
import Navbar from '@/app/components/Navbar/Navbar';

interface File {
  ID: string;
  UserID: string;
  FileName: string;
  S3Bucket: string;
  S3ObjectKey: string;
  CreatedAt: string;
  IsPublic: boolean;
  Size: number;
  ContentType: string;
}

async function listFiles(): Promise<File[]> {
  const url = `${process.env.AUTH_BACKEND_URL}/api/v1/filelist`;

  let accessToken = await getAccessToken();

  const resp = await fetch(url, {
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    },
  });

  if (resp.ok) {
    const data = await resp.json();
    return data as File[];
  }

  throw new Error("Failed to fetch data. Status: " + resp.status);
}
export default async function Dashboard() {
  const session = await getServerSession(authOptions);

  if (session) {
    try {
      const files = await listFiles();          
    return (
    <main>
      <Navbar />
      <FileTree filenames={files.map(file => file.FileName)} />
    </main>
    );
  } catch (err) {
    console.error(err);

      return (
        <main>
          <h1 className="text-4xl text-center">Your Files</h1>
          <p className="text-primary text-center text-lg">
            You Don&apos;t Have any files yet!
          </p>
        </main>
      );
    }
  }

  redirect("/unauthorised");
}