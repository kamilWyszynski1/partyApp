import { Component, OnInit, NgZone } from '@angular/core';
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
  admin_login = 'front@wp.pl';
  admin_password = 'front'

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
    private router: Router,
    private zone: NgZone
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
    body.set('client_secret', this.CLIENT_SECRET);

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
        this.zone.run(() => this.router.navigate(['/party']));
      },
      err => {
        console.log(err);
        this.snackbar.open('Invalid email or password', 'OK!');
      }
    )
  }

  signUp(){
    let email = this.signUpForm.value['email'];
    let pass = this.signUpForm.value['password'];
    let pass2 = this.signUpForm.value['repeatPassword'];

    if (pass !== pass2){
      this.snackbar.open('Passwords must be identical', 'OK!');
      return;
    }

    let body = new URLSearchParams();
    body.set('grant_type', 'password');
    body.set('username', this.admin_login);
    body.set('password', this.admin_password);
    body.set('client_id', this.CLIENT_ID);
    body.set('client_secret', this.CLIENT_SECRET);

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
        let token = res['access_token'];
        let create_options = {
          headers : new HttpHeaders({
            'Authorization': 'Bearer '+token
          })
        }
        let create_url = 'http://127.0.0.1:8000/clients/';
        this.http.post(
          create_url, {
            'email': email,
            'password': pass
          },
          create_options
        ).subscribe(
          res => {
            this.signUpView();
            this.snackbar.open('Account Created!', 'OK!');
          },
          err => {
            console.log(err);
            this.snackbar.open('Email is taken!', 'OK!');
          }
        )
      },
      err => {
        console.log('err1', err);
      }
    )
  }

  signUpView(){
    this.loginView = !this.loginView;

  }
}
