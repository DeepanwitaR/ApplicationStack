import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {LeftSideModule} from './left-side/left-side.module'
import { AppComponent } from './app.component';
import {RouterModule} from '@angular/router';  
import { Left1Component } from './left-side/left1/left1.component';
import { Left2Component } from './left-side/left2/left2.component';
import { Service1Service } from './service1.service';
import {HttpserviceService} from './httpservice.service';
import {HttpClientModule} from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule, 
    LeftSideModule,
    HttpClientModule, // We are also putting in the httpServices into the provider list automatically.
    RouterModule.forRoot([
      { 
        path: '', 
        component: AppComponent
      },
      { 
        path: 'left1', 
        component: Left1Component
      }, //path is and their corresponding component.
      {
        path: 'left2', 
        component: Left2Component
      },
      { 
        path: '**', 
        component: Left1Component
      }
      // {
      //   path: 'left2/:parameter_dynamicallyput', 
      //   component: Left2Component
      // }
      // {
    ]) //root routes, where we hit the //we must have a routerOutlet, where to display the corresponding comp.
  ],
  providers: [
    Service1Service, 
    HttpserviceService
  ],
  bootstrap: [
    AppComponent
  ]
})
export class AppModule { }
