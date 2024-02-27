import React from 'react';

const FileItem = ({ children }) => <li>{children}</li>;

const FileDetails = ({ open, label, children }) => (
  <details open={open}>
    <summary>{label}</summary>
    <ul>{children}</ul>
  </details>
);

const FileTree = ({ filenames }) => (
    <ul>
      {filenames.map((filename, index) => (
        <FileItem key={index}>
          <a>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth="1.5"
              stroke="currentColor"
              className="w-4 h-4"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m2.25 0H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z"
              />
            </svg>
            {filename}
          </a>
          <FileDetails open={false} label={`Details for ${filename}`}>
            {/* Additional details for each file can be placed here */}
          </FileDetails>
        </FileItem>
      ))}
    </ul>
  );
  

export default FileTree;