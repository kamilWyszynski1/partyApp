import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { AppComponent } from './app.component';
import { UserComponent } from './user/user.component';

const appRoutes: Routes = [
  {path: '**', component: (() => {
    return localStorage.getItem('user') ? HomeComponent : UserComponent;
    })()
  },
  {
    path: 'party', component: HomeComponent
  },
];

@NgModule({
  imports: [RouterModule.forRoot(
    appRoutes,
    { enableTracing: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
