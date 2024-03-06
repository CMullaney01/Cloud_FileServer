"use client";
import React, { useState, useEffect } from "react";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";

const StreamVideoButton = ({ fileName }) => {
    const { data: session, status } = useSession();
    const router = useRouter();

    useEffect(() => {
        if (status === "unauthenticated") {
            router.push("/unauthorized");
        }
    }, [status, router]);

    const [errorMsg, setErrorMsg] = useState(null);

    const handleStream = async () => {
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
    
            // Create a temporary video element
            const video = document.createElement('video');
            video.controls = true; // Show controls for the video player
            video.style.width = '100%'; // Set width to fill container
            video.style.height = 'auto'; // Maintain aspect ratio
    
            // Set the source of the video to the presigned URL
            video.src = presignedURL;
    
            // Append the video element to the DOM
            document.body.appendChild(video);
    
            // Play the video
            video.play();
    
            router.push("/Dashboard"); // Optionally, redirect the user after initiating the stream
        } catch (error) {
            setErrorMsg("Failed to stream video: " + error.message);
        }
    };

    return (
        <div>
            <button onClick={handleStream}>Stream {fileName}</button>
            {errorMsg && <div>{errorMsg}</div>}
        </div>
    );
};

export default StreamVideoButton;
