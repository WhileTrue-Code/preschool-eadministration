import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { AppointmentsComponent } from './components/appointment/appointments/appointments.component';
import { CompetitionsComponent } from './components/competitions/competitions/competitions.component';
import { CompetitionViewComponent } from './components/competitions/competition-view/competition-view.component';
import { RegisterAprComponent } from './components/register-apr/register-apr.component';
import { MyAprsComponent } from './components/my-aprs/my-aprs.component';
import { CompetitionAddComponent } from './components/competitions/competition-add/competition-add.component';

const routes: Routes = [
  {
    path: "Login",
    component: LoginComponent
  },
  {
    path: "Register",
    component: RegisterComponent
  },
  {
    path: "Appointments",
    component: AppointmentsComponent
  },
  {
    path: "Competitions",
    component: CompetitionsComponent
  },
  {
    path: "Competition-View/:id",
    component: CompetitionViewComponent
  },
  {
    path: "Competition-Add",
    component: CompetitionAddComponent
  },
  {
    path: "RegisterApr",
    component: RegisterAprComponent
  },
  {
    path: "MyAprs",
    component: MyAprsComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
