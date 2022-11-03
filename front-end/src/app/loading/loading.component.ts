import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { UploadService } from '../_services/upload.service';

@Component({
  selector: 'app-loading',
  templateUrl: './loading.component.html',
  styleUrls: ['./loading.component.scss']
})
export class LoadingComponent implements OnInit, OnDestroy {

  ready: boolean = false;
  error: boolean = false;
  errorMessage: string = '';
  decrypting: boolean = true;
  sub: Subscription;
  token: string;
  uuid: string;

  constructor(private route: ActivatedRoute, private encryptService: UploadService, private router: Router) {
    this.sub = this.route.params.subscribe((params) => {
      this.token = params['token'];
    });
  }

  ngOnInit(): void {
    this.decryptFiles();
  }

  ngOnDestroy(): void {
    this.sub.unsubscribe();
  }

  decryptFiles(): void {
    setTimeout(() => {
      this.encryptService.decryptFiles(this.token).subscribe((response) => {
        if (response.status) {
          this.decrypting = false;
          this.ready = true;
          this.uuid = response.uuid;
        }
      }, (error) => {
        this.decrypting = false;
        this.error = true;
        this.errorMessage = error.message;
      })
    }, 5000);
  }

  downloadFiles(): void {
    this.encryptService.downloadFiles(this.uuid).subscribe((file) => {
      const blob = new Blob([file], {
        type: 'application/zip'
      });
      const url = window.URL.createObjectURL(blob);
      window.open(url);
      setTimeout(() => {
        this.router.navigate(['/home']);
      }, 2000);
    });
  }

}
