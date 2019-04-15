import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { MatSnackBar } from '@angular/material';
import { LocalStorageService } from 'angular-web-storage';
import { Router } from '@angular/router';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  loginView: boolean = true;
  CLIENT_ID = '2N1ufWHdDYhWTA7bFgLqrNdFkQc7WiKfmq7mZlIX';
  CLIENT_SECRET = 'MEMbMT7UBCQLhnYo2sQit8JyQ8qYzIxdhy61wwWRk63lIlHCp2bjRKL5zsCfQFRm0KUZB95ZdjhETeQD1SNXT8wzidMM266lxxEQ49GunKvDdI2DWJGvJBydq9r39axA';


  userForm = new FormGroup({
    email: new FormControl(''),
    password: new FormControl('')
  });

  signUpForm = new FormGroup({
    email: new FormControl(''),
    password: new FormControl(''),
    repeatPassword: new FormControl('')
  });


  constructor(
    private http: HttpClient,
    private snackbar: MatSnackBar,
    private local: LocalStorageService,
    private router: Router
  ) { }

  ngOnInit() {
  }

  signIn(){
    let email = this.userForm.value['email'];
    let password = this.userForm.value['password'];
    let body = new URLSearchParams();
    body.set('grant_type', 'password');
    body.set('username', email);
    body.set('password', password);
    body.set('client_id', this.CLIENT_ID);
    body.set('client_secret', this.CLIENT_SECRET)

    let options = {
      headers: new HttpHeaders({
        'Content-Type': 'application/x-www-form-urlencoded'
      })
    };

    let url = 'http://127.0.0.1:8000/o/token/';
    
    this.http.post(
      url, body.toString(), options
    ).subscribe(
      res => {
        console.log(res['access_token']);
        this.local.set('token', res['access_token']);
        this.router.navigate(['/party?refresh=1']);
      },
      err => {
        console.log(err);
        this.snackbar.open('Invalid email or password', 'OK!');
      }
    )
  }

  signUpView(){
    this.loginView = !this.loginView;

  }
}
