import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {Left1Component} from './left-side/left1/left1.component'
import {Left2Component} from './left-side/left2/left2.component'
import {LeftSideModule} from './left-side/left-side.module'
import { AppComponent } from './app.component';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule
    // , LeftSideModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
