import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { AppComponent } from './app.component';
import { UserComponent } from './user/user.component';
import { LayoutComponent } from './layout/layout.component';

const appRoutes: Routes = [
  {
    path: 'login', component: UserComponent
  },
  {
    path: '', component: LayoutComponent,
    children:[
      {
        path: 'party',
        component: HomeComponent
      }
    ]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(
    appRoutes,
    { enableTracing: false })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
