import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {LeftSideModule} from './left-side/left-side.module'
import { AppComponent } from './app.component';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule, LeftSideModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
