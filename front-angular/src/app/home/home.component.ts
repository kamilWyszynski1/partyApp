import { Component, OnInit, ViewChild, Input } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { FormGroup, FormControl } from '@angular/forms';
import {} from 'googlemaps';
import { LocalStorageService } from 'angular-web-storage';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
 
  /*
    Creation of Party and Order models using
    Party/Order Service and its API.

    Firstly we create Party object, get its id 
    and pass it into Order creation which returns its
    id, then we update Party's foreing key of Order 
    and its done.

    PARTY
    {
      'user': ,
      'location': ,
      'hostesses' :
    }

    ORDER
    {
      'party_id' :,
      'content': 
    }

  */
 scene = 0;
 constructor(
   private http: HttpClient,
   public local: LocalStorageService,
   
   ){}

 KEY = 'value';
 value: any = null;
 
 party_id : number;
 

 @ViewChild('map') mapElement: any;
 map: google.maps.Map;

 orderForm = new FormGroup({
   content: new FormControl(''),
   user: new FormControl(''),
   location: new FormControl(''),
   hostess: new FormControl('')
 });

 httpOptions ={
   headers: new HttpHeaders({
     'Content-Type': 'application/json'
   })
 };
 
 onSubmit(){
   this.http.post(
     'http://0.0.0.0:5000/party', 
     {
       "user": 1,
       "location":this.orderForm.value['location'],
       "order": "",
       "hostesses": this.orderForm.value['hostess']
     },
     this.httpOptions
   )
   .subscribe(
     res => {
       console.log(res['id']);
       this.http.post(
         'http://0.0.0.0:3000/order',
         {
           'party': res['id'],
           'content': this.orderForm.value['content']
         },
         this.httpOptions
       )
       .subscribe(
         res => {
           this.http.put(
             'http://0.0.0.0:5000/party/' + res['party'],
             {
               'order': res['_id']
             },
             this.httpOptions
           ).subscribe(
             res => {
               console.log('yea!');
             },
             err => {
               console.log(err);
             }
           )
         },
         err => {
           console.log(err);
         }
       )
     },
     err => {
       console.log(err);
     }
   );
 }

  initMap(){}

  ngOnInit() {
    setTimeout(() => {
      const mapProperties = {
        center: new google.maps.LatLng(52.229675, 21.012230),
        zoom: 15,
        mapTypeId: google.maps.MapTypeId.ROADMAP
    };
    this.map = new google.maps.Map(this.mapElement.nativeElement, mapProperties);
    }, 1000);
  }

  session(){
    localStorage.clear();
  }

  open(){
    console.log('asd');
  }

  changeView(){
    this.scene += 1;
  }
}
