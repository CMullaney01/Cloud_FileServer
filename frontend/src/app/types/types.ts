export interface File {
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

export interface UploadFileResponse {
    file: File;
    presignedURL: string;
}