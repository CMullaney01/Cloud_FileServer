"use client";
import React, { useState, useEffect } from "react";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";

const DownloadFileButton = ({ fileName }) => {
    const { data: session, status } = useSession();
    const router = useRouter();

    useEffect(() => {
        if (status === "unauthenticated") {
            router.push("/unauthorized");
        }
    }, [status, router]);

    const [errorMsg, setErrorMsg] = useState(null);

    const handleDownload = async () => {
        try {
            const url = `/api/files?filename=${encodeURIComponent(fileName)}`;
            const resp = await fetch(url, { method: "GET" });
            if (!resp.ok) {
                const json = await resp.json();
                setErrorMsg("Unable to get download URL: " + json.error);
                return;
            }
    
            const data = await resp.json();
            const presignedURL = data.data;
    
            // Fetch the file using the presigned URL
            const downloadResp = await fetch(presignedURL, { method: "GET" });
    
            if (!downloadResp.ok) {
                setErrorMsg("Download failed: " + downloadResp.statusText);
                return;
            }
    
            // Convert the blob response into a downloadable file
            const blob = await downloadResp.blob();
            const blobURL = URL.createObjectURL(blob);
    
            // Create a temporary anchor element
            const anchor = document.createElement('a');
            anchor.href = blobURL;
            anchor.download = fileName; // Set the download attribute to specify the filename
    
            // Programmatically click on the anchor to trigger the download
            anchor.dispatchEvent(new MouseEvent('click'));
    
            // Clean up by revoking the blob URL
            URL.revokeObjectURL(blobURL);
    
            router.push("/dashboard"); // Optionally, redirect the user after initiating the download
        } catch (error) {
            setErrorMsg("Failed to get download URL: " + error.message);
        }
    };

    return (
        <div>
            <button onClick={handleDownload}>Download {fileName}</button>
            {errorMsg && <div>{errorMsg}</div>}
        </div>
    );
};

export default DownloadFileButton;
