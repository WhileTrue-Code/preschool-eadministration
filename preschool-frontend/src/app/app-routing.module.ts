import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { AppointmentsComponent } from './components/appointment/appointments/appointments.component';
import { CompetitionsComponent } from './components/competitions/competitions/competitions.component';
import { CompetitionViewComponent } from './components/competitions/competition-view/competition-view.component';
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
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
