import React from 'react'
import architectureImage from '../../../../public/Cloudify.png';
import Image from 'next/image';

const ArchitectureCard = () => {
  return (
    <div className="card w-50 bg-base-100 shadow-xl hover-scale">
      <figure> <Image  className='bg-primary'style={{ borderRadius: '2.5rem'}}src={architectureImage} alt="Your Image" /></figure>
      <div className="card-body">
          <h2 className="card-title">
          Documentation
          <div className="badge badge-secondary">updated</div>
            </h2>
            <p><strong>Would you like to integrate your backend REST API with our project? Follow these steps to set up and connect to your own S3 instance:</strong></p>
        <ul className="ml-4 list-disc">
          <li><strong>Check out the documentation:</strong> Refer to the provided documentation to learn how to connect your backend REST API with our project.</li>
          <li><strong>Set up your own S3 instance:</strong> Set up your own Amazon S3 (Simple Storage Service) instance or configure your own S3-compatible storage solution.</li>
          <li><strong>Interact:</strong> Now you can interact with your files through our frontend!</li>
        </ul>
        <p>By following these steps, you&apos;ll be able to seamlessly integrate your backend REST API and S3 instance with our project.</p>
      </div>
      <div className="card-actions justify-end mb-1 mr-1">
      <div className="badge badge-outline">Documentation</div> 
      <div className="badge badge-outline">Personal storage solution</div>
      </div>
    </div>
  )
}

export default ArchitectureCard