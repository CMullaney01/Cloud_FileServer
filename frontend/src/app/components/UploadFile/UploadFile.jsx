"use client"
import React, { useState, useEffect, useRef } from "react";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";

const UploadFile = () => {
  const { data: session, status } = useSession();
  const router = useRouter();

  useEffect(() => {
    if (status === "unauthenticated") {
      router.push("/unauthorized");
      router.refresh();
    }
  }, [session, status, router]);

  const fileInputRef = useRef();
  const [errorMsg, setErrorMsg] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();

    const selectedFile = fileInputRef.current.files[0];

    if (!selectedFile) {
      setErrorMsg("Please select a file.");
      return;
    }

    const postBody = {
      FileName: selectedFile.name,
      ContentType: selectedFile.type,
      size: selectedFile.size,
    };

    try {
      const resp = await fetch("/api/files", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(postBody),
      });

      if (resp.ok) {
        router.push("/dashboard");
        router.refresh();
      } else {
        const json = await resp.json();
        setErrorMsg("Unable to call the API: " + json.error);
      }
    } catch (err) {
      setErrorMsg("Unable to call the API: " + err);
    }
  };

  return (
    <main>
      <h1 className="text-4xl text-center">Upload File</h1>
      <input
        type="file"
        className="file-input file-input-bordered file-input-primary w-full max-w-xs"
        ref={fileInputRef}
      />
      <button type="submit" onClick={handleSubmit}>
        Upload
      </button>
      {errorMsg && <div>{errorMsg}</div>}
    </main>
  );
};

export default UploadFile;
