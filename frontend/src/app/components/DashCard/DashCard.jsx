import React from 'react'
import dashImage from '../../../../public/Dashboard.png';
import Image from 'next/image';


const DashCard = () => {
  return (
    <div className="card w-50 bg-base-100 shadow-xl hover-scale">
      <figure> <Image  className='bg-primary'style={{ borderRadius: '2.5rem'}}src={dashImage} alt="Your Image" /></figure>
      <div className="card-body">
          <h2 className="card-title">
          Your File Dashboard
          <div className="badge badge-secondary">updated</div>
            </h2>
            <p><strong>Login, Create an Account and browse your files!</strong></p>
        <ul className="ml-4 list-disc">
          <li><strong>Manage your files:</strong> Upload, download, and organize your files effortlessly.</li>
          <li>
            <strong>Stream your videos:</strong> Enjoy seamless streaming of your videos directly from the website.
        </li>
        </ul>
        <p>By utilizing these features, you can efficiently manage and navigate through your files with ease.</p>
      </div>
      <div className="card-actions justify-end mb-1 mr-1">
      <div className="badge badge-outline">File Managing</div> 
      <div className="badge badge-outline">Video Streaming</div>
      </div>
    </div>
  )
}

export default DashCard