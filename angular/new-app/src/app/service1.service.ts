import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class Service1Service {

  constructor() { }
  
  getDetails() {
    return [
      {"id" : "1", "name" : "one"},
      {"id" : "2", "name" : "two"}
    ];
  }
}
