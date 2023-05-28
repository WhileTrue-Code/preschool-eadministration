import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { AppointmentsDoctorComponent } from './components/appointment/appointments-doctor/appointments.component';
import { CompetitionsComponent } from './components/competitions/competitions/competitions.component';
import { CompetitionViewComponent } from './components/competitions/competition-view/competition-view.component';
import { RegisterAprComponent } from './components/register-apr/register-apr.component';
import { MyAprsComponent } from './components/my-aprs/my-aprs.component';
import { CompetitionAddComponent } from './components/competitions/competition-add/competition-add.component';
import { WelcomeComponent } from './components/welcome/welcome.component';
import { AppointmentAddComponent } from './components/appointment/appointment-add/appointment-add.component';
import { AppointmentViewComponent } from './components/appointment/appointment-view/appointment-view.component';
import { AppointmentsRegularComponent } from './components/appointment/appointments-regular/appointments-regular.component';
import { VaccinationsDoctorComponent } from './components/vaccination/vaccinations-doctor/vaccinations-doctor.component';
import { VaccinationsRegularComponent } from './components/vaccination/vaccinations-regular/vaccinations-regular.component';
import { VaccinationAddComponent } from './components/vaccination/vaccination-add/vaccination-add.component';
import { VaccinationViewComponent } from './components/vaccination/vaccination-view/vaccination-view.component';
import {MarriageComponent} from "./components/marriage/marriage.component";
import {ChooseServiceComponent} from "./components/choose-service/choose-service.component";

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
    path: "Appointments-Doctor",
    component: AppointmentsDoctorComponent
  },
  {
    path: "Appointments-Regular",
    component: AppointmentsRegularComponent
  },
  {
    path: "Appointment-Add",
    component: AppointmentAddComponent
  },
  {
    path: "Appointment-View/:id",
    component: AppointmentViewComponent
  },
  {
    path: "Vaccinations-Doctor",
    component: VaccinationsDoctorComponent
  },
  {
    path: "Vaccinations-Regular",
    component: VaccinationsRegularComponent
  },
  {
    path: "Vaccination-Add",
    component: VaccinationAddComponent
  },
  {
    path: "Vaccination-View/:id",
    component: VaccinationViewComponent
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
  },
  {
    path: "Welcome",
    component: WelcomeComponent
  },
  {
    path: "Marriage",
    component: MarriageComponent
  },
  {
    path: "pre-login",
    component: ChooseServiceComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
