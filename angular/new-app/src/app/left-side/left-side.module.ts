import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Left1Component } from './left1/left1.component';
import { Left2Component } from './left2/left2.component';

@NgModule({
  declarations: [
     Left1Component, Left2Component
  ],
  imports: [
    CommonModule
  ],
  exports: [
    CommonModule
    ,Left2Component, Left1Component
  ]
})
export class LeftSideModule { }
