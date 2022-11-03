export interface UploadResponse {
    status: boolean;
    message: string;
    token?: string;
    error?: string;
}

export interface DecryptResponse {
    status: boolean;
    uuid: string;
    message: string;
}