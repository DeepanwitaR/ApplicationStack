import { Component } from '@angular/core';
import { Service1Service } from './service1.service';
import { HttpserviceService } from './httpservice.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'new-app';
  // list = [];
  //service to be accessed via dependency injection.
  constructor(private list1: Service1Service, private httpList: HttpserviceService) {
    console.log(list1.getDetails());
    console.log(httpList.getHttpDetails());
  }
}
