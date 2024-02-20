# Cloud_FileServer
Extending the File Server project written in Go to make more use of the cloud. Please read the roadmap below, for more information on the design choices please look at the research folder.


# Project Roadmap

<font color="green">
## ✅ Completed: Milestone 1: Authentication and Login
</font>
- Identified keycloak as an appropriate auth service. 
- Set up a user database to store user credentials securely.
- Develop a login page and authentication API endpoints.
- Able to implement access control here with keycloak by protecting end points as such we have removed that milestone
If you are looking to get started on keycloak yourself, I folowed a great tutorial https://www.youtube.com/watch?v=1u8GlfKyB_Q&t=810s.

### current setup
![Example Image](research/images/auth-login.png)

<font color="orange">
## 🚀 In Progress: File Upload and Storage:
</font>
- Allow users to upload files to the server.
- upload meta data to the maria db which will hold file url and sign the URL before passing it back to the client.
- Integrate cloud storage on Amazon S3 service 
- Make it look good!

## File Management:
- Develop APIs for file management operations such as create, read, update, and delete (CRUD) operations on files.
- Implement file organization features like folders or directories.

## File Sharing:
- Enable users to share files or folders with other users.
- Implement secure sharing mechanisms with options for setting permissions and expiration dates.

## Versioning and Revision History:
- Implement version control for files to keep track of changes over time.
- Allow users to revert to previous versions if needed.

## Search:
- Develop search functionality to allow users to search for files based on metadata or content.


## Video Streaming:
- Integrate video streaming capabilities to allow users to stream video files stored on the server.
- Implement adaptive bitrate streaming for better performance across different network conditions.

## Content Delivery Network (CDN):
- Utilize a CDN to cache and deliver static assets like videos, improving performance and scalability.

## Transcoding and Encoding:
- Implement video transcoding and encoding services to support multiple formats and resolutions for streaming.

## Analytics and Monitoring:
- Set up monitoring and analytics tools to track server performance, user activity, and usage patterns.
- Implement logging to capture and analyze events for troubleshooting and optimization.

## Security Enhancements:
- Implement encryption for data at rest and in transit to ensure data security.
- Set up intrusion detection and prevention systems (IDS/IPS) to detect and mitigate security threats.

## High Availability and Scalability:
- Design the architecture for high availability with redundancy and failover mechanisms.
- Implement auto-scaling to handle varying loads and traffic spikes efficiently.

## Backup and Disaster Recovery:
- Set up regular backups of data stored on the server.
- Develop a disaster recovery plan and implement mechanisms for data restoration in case of failures.

## Compliance and Regulation:
- Ensure compliance with relevant regulations and standards such as GDPR, HIPAA, or PCI DSS.
- Implement features for data retention policies and compliance reporting.