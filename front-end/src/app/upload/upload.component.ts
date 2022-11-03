import { Component, OnInit, ViewChild } from '@angular/core';
import { MessageService } from 'primeng/api';
import { FileUpload } from 'primeng/fileupload';
import { UploadService } from '../_services/upload.service';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.scss']
})
export class UploadComponent implements OnInit {

  uploadedFiles: any[] = [];
  multiple = true;
  maxSize = 10000000000;
  token: string;
  loading = false;
  displayPosition: boolean;
  @ViewChild('upload') fileUpload: FileUpload;

  position: string;
  constructor(private uploadService: UploadService, private msgService: MessageService) { }

  ngOnInit(): void {
  }

  onUpload(event: any): void {
    console.log(event);
  }

  selectFile(event: any): void {
    this.uploadedFiles = event.currentFiles;
  }

  uploadFiles(event: any): void {
    this.loading = true;
    const formdata: FormData = new FormData();
    event.files.forEach((file: File) => {
      formdata.append('multiplefiles', file);
    })
    this.uploadService.uploadFiles(formdata).subscribe((response) => {
      this.token = response.token ? response.token : '';
      this.msgService.add({ severity: 'success', summary: 'Upload', detail: response.message });
      setTimeout(() => {
        this.showToken('bottom-right');
      }, 800);
      this.loading = false;
    }, (error) => {
      this.loading = false;
      this.msgService.add({ severity: 'error', summary: 'Upload', detail: error.message });
    });
  }

  showToken(position: string) {
    this.position = position;
    this.fileUpload.clear();
    this.uploadedFiles = [];
    this.displayPosition = true;
  }


  copySharingLink(): void {
    this.msgService.add({ severity: 'info', summary: 'Clipboard', detail: 'Link copied to clipboard' });
    this.displayPosition = false;
  }

}
