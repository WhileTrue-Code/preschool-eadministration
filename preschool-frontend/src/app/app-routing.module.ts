import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { AppointmentsDoctorComponent } from './components/eZdravstvo/appointment/appointments-doctor/appointments.component';
import { CompetitionsComponent } from './components/competitions/competitions/competitions.component';
import { CompetitionViewComponent } from './components/competitions/competition-view/competition-view.component';
import { RegisterAprComponent } from './components/register-apr/register-apr.component';
import { MyAprsComponent } from './components/my-aprs/my-aprs.component';
import { CompetitionAddComponent } from './components/competitions/competition-add/competition-add.component';
import { WelcomeComponent } from './components/welcome/welcome.component';
import { AppointmentAddComponent } from './components/eZdravstvo/appointment/appointment-add/appointment-add.component';
import { AppointmentViewComponent } from './components/eZdravstvo/appointment/appointment-view/appointment-view.component';
import { AppointmentsRegularComponent } from './components/eZdravstvo/appointment/appointments-regular/appointments-regular.component';
import { VaccinationsDoctorComponent } from './components/eZdravstvo/vaccination/vaccinations-doctor/vaccinations-doctor.component';
import { VaccinationsRegularComponent } from './components/eZdravstvo/vaccination/vaccinations-regular/vaccinations-regular.component'
import { VaccinationAddComponent } from './components/eZdravstvo/vaccination/vaccination-add/vaccination-add.component';
import { VaccinationViewComponent } from './components/eZdravstvo/vaccination/vaccination-view/vaccination-view.component';
import {MarriageComponent} from "./components/marriage/marriage.component";
import {ChooseServiceComponent} from "./components/choose-service/choose-service.component";
import {RegularOrAdminComponent} from "./components/regular-or-admin/regular-or-admin.component";
import { VrticPocetnaComponent } from './components/preschool/vrtic-pocetna/vrtic-pocetna.component';
import { VrticAddComponent } from './components/preschool/vrtic-add/vrtic-add.component';
import { VrticViewComponent } from './components/preschool/vrtic-view/vrtic-view.component';
import { PrijavaComponent } from './components/preschool/prijava/prijava.component';
import { PrijavaMainComponent } from './components/preschool/prijava-main/prijava-main.component';
import { AddPersonRegistryComponent } from './components/eZdravstvo/add-person-registry/add-person-registry.component'; 
import { RegisterCrosoComponent } from './components/register-croso/register-croso.component';
import { MyCrososComponent } from './components/my-crosos/my-crosos.component';
import { RegisterEmployeeComponent } from './components/register-employee/register-employee.component';
import { CompanyEmployeesComponent } from './components/company-employees/company-employees.component';
import {ViewMyRegistryComponent} from "./components/view-my-registry/view-my-registry.component";
import { ZdravstvenoStanjeAddComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstveno-stanje-add/zdravstveno-stanje-add.component';
import { ZdravstvenaStanjaDoctorComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstvena-stanja-doctor/zdravstvena-stanja-doctor.component';
import { ZdravstvenoStanjeViewMyComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstveno-stanje-view-my/zdravstveno-stanje-view-my.component';
import { VaccinationsMyRegularComponent } from './components/eZdravstvo/vaccination/vaccinations-my-regular/vaccinations-my-regular.component';
import { UserDiedComponent } from './components/eZdravstvo/user-died/user-died.component';

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
    path: "New-Person",
    component: AddPersonRegistryComponent
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
    path: "Zdravstvena-Stanja-Doctor",
    component: ZdravstvenaStanjaDoctorComponent
  },
  {
    path: "Zdravstveno-Stanje-View-My",
    component: ZdravstvenoStanjeViewMyComponent
  },
  {
    path: "Vaccinations-My-Regular",
    component: VaccinationsMyRegularComponent
  },
  {
    path: "Zdravstveno-Stanje-Add",
    component: ZdravstvenoStanjeAddComponent
  },
  {
    path: "User-Died",
    component: UserDiedComponent
  },
  {
    path: "Competitions",
    component: CompetitionsComponent
  },
  {
    path: "PocetnaVrtic",
    component: VrticPocetnaComponent
  },
  {
    path: "Vrtic-Add",
    component: VrticAddComponent
  },
  {
    path: "Vrtic-View/:id",
    component: VrticViewComponent
  },
  {
    path: "Competition-View/:id",
    component: CompetitionViewComponent
  },
  {
    path: "Prijava/:id",
    component: PrijavaComponent
  },
  {
    path: "PregledPrijava/:id",
    component: PrijavaMainComponent
  },
  {
    path: "Competition-Add/:id",
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
    path: "chose-service",
    component: ChooseServiceComponent
  },
  {
    path: "regular-or-admin",
    component: RegularOrAdminComponent
  },
  {
    path:"RegisterCroso",
    component: RegisterCrosoComponent
  },
  {
    path: "MyCrosos",
    component: MyCrososComponent
  },
  {
    path: "RegisterEmployee",
    component: RegisterEmployeeComponent
  },
  {
    path: "CompanyEmployees",
    component: CompanyEmployeesComponent
  },{
    path: "view-my-registry",
    component: ViewMyRegistryComponent
  }

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
