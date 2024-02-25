import React from 'react'
import architectureImage from '../../../../public/Cloudify.png';
import Image from 'next/image';


const TutorialCard = () => {
  return (
    <div className="card w-50 bg-base-100 shadow-xl hover-scale">
    <figure> <Image  className='bg-primary'style={{ borderRadius: '2.5rem'}}src={architectureImage} alt="Your Image" /></figure>
    <div className="card-body">
        <h2 className="card-title">
        Documentation
        <div className="badge badge-secondary">updated</div>
          </h2>
          <p><strong>Want to know how to bild the project from scratch? Enjoy this step by step tutorial</strong></p>
      <ul className="ml-4 list-disc">
        <li><strong>Set up Authentication:</strong>Learn the basics of keycloak to secure authentication for your applications</li>
        <li><strong>Build your own Rest Api</strong>What is a rest api how does it work , follow along and find out</li>
        <li><strong>Integrate with the front end</strong> Learn how your application can communicate with a rest api!</li>
      </ul>
      <p>By following these steps, you&apos;ll be able to create your own file server and learn valuable skills along the way</p>
    </div>
    <div className="card-actions justify-end mb-1 mr-1">
    <div className="badge badge-outline">Pro developer</div> 
    </div>
  </div>
  )
}

export default TutorialCard