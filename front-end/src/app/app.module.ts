import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { ButtonModule } from 'primeng/button';
import { HowComponent } from './how/how.component';
import { UploadComponent } from './upload/upload.component';
import { FileUploadModule } from 'primeng/fileupload';
import { HttpClientModule } from '@angular/common/http';
import { HomeComponent } from './home/home.component';
import { MessagesModule } from 'primeng/messages';
import { MessageModule } from 'primeng/message';
import { MessageService } from 'primeng/api';
import { ToastModule } from 'primeng/toast';
import { ProgressBarModule } from 'primeng/progressbar';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { ClipboardModule } from '@angular/cdk/clipboard';
import { LoadingComponent } from './loading/loading.component';


@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    HowComponent,
    UploadComponent,
    HomeComponent,
    LoadingComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    ButtonModule,
    FileUploadModule,
    HttpClientModule,
    MessageModule,
    MessagesModule,
    ToastModule,
    ProgressBarModule,
    DialogModule,
    InputTextModule,
    ClipboardModule
  ],
  providers: [MessageService],
  bootstrap: [AppComponent],
  exports: [AppRoutingModule, AppComponent]
})
export class AppModule { }
