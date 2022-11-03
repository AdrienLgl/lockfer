import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { DecryptResponse, UploadResponse } from '../_models/file';

@Injectable({
  providedIn: 'root'
})
export class UploadService {

  constructor(private http: HttpClient) { }

  uploadFiles(formdata: any): Observable<UploadResponse> {
    return this.http.post<UploadResponse>('/api/v1/upload', formdata);
  }

  decryptFiles(token: string): Observable<DecryptResponse> {
    const body = { token };
    return this.http.post<DecryptResponse>('/api/v1/decrypt', body);
  }

  downloadFiles(uuid: string): Observable<any> {
    const headers = { 'responseType': 'blob' as 'json' };
    const body = { title: 'Test' };
    return this.http.post<any>('/api/v1/download/' + uuid, body, headers);
  }
}
