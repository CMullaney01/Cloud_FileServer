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
    
      if (!resp.ok) {
        const json = await resp.json();
        setErrorMsg("Unable to call the API: " + json.error);
        return;
      }
    
      const data = await resp.json();
      const presignedURL = data.data.presignedURL.presignedURL;
      // console.log("presignedURL:", presignedURL)
      // Upload file to S3 using presigned URL
      const uploadResp = await fetch(presignedURL, {
        method: "PUT",
        body: selectedFile,
        headers: {
          "Content-Type": selectedFile.type,
        },
      });
    
      if (!uploadResp.ok) {
        setErrorMsg("Upload to S3 failed: " + uploadResp.statusText);
        return;
      }
    
      // If upload to S3 is successful, confirm the upload
      const confirmResp = await fetch("/api/files/confirm", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(postBody), // Adjust confirmBody if needed
      });
    
      if (!confirmResp.ok) {
        setErrorMsg("Upload confirmation failed: " + confirmResp.statusText);
      } else {
        // If everything is successful, navigate to the dashboard
        router.push("/Dashboard");
        router.refresh();
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
