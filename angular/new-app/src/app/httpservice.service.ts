import { Injectable, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class HttpserviceService {

  constructor(private httpClient: HttpClient) { } //to send the get.
  
  ngOnInit(){
    //inits right after the constructor.
  }
  getHttpDetails() {
    //get request to server, and then returns an observable.
    // https://angular.io/guide/observables-in-angular 
    // Observable: the stream of data coming in, masked by an obj
    // Observer: a class that has methods, to use up the observable. next(), complete(), error(), ...

    return this.httpClient.get("./assets/javascriptFile.js");
  }  
}

