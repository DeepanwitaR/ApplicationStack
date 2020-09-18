import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Left1Component } from './left1/left1.component';
import { Left2Component } from './left2/left2.component';

@NgModule({
  declarations: [Left1Component, Left2Component],
  imports: [
    CommonModule
    // , Left1Component, Left2Component
  ],
  exports: [
    CommonModule
    ,Left1Component,LeftSideModule
    // Left2Component, LeftSideModule
  ]
})
export class LeftSideModule { }
