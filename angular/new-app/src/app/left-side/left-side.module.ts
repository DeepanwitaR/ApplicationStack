import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Left1Component } from './left1/left1.component';
import { Left2Component } from './left2/left2.component';
import {RouterModule} from '@angular/router';
@NgModule({
  declarations: [
     Left1Component, Left2Component
  ],
  imports: [
    CommonModule,
    RouterModule.forRoot([
      { 
        path: '', 
        component: Left1Component
      }, //path is and their corresponding component.
      {
        path: 'left2', 
        component: Left2Component
      }
      // {
      //   path: 'left2/:parameter_dynamicallyput', 
      //   component: Left2Component
      // }
      // {
      //   path: '**', 
      //   component: NotFoundComponent
      // } //wildcard
    ]) //root routes, where we hit the //we must have a routerOutlet, where to display the corresponding comp.
  ],
  exports: [
    CommonModule
    ,Left2Component, Left1Component
  ]
})
export class LeftSideModule { }
