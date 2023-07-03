import { NgModule } from '@angular/core';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { MatToolbarModule } from '@angular/material/toolbar';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { AuthInterceptor } from './services/auth.interceptor';
import { AppointmentAddComponent } from './components/eZdravstvo/appointment/appointment-add/appointment-add.component';
import { AppointmentItemComponent } from './components/eZdravstvo/appointment/appointment-item/appointment-item.component';
import { AppointmentListComponent } from './components/eZdravstvo/appointment/appointment-list/appointment-list.component';
import { AppointmentViewComponent } from './components/eZdravstvo/appointment/appointment-view/appointment-view.component';
import { AppointmentsDoctorComponent } from './components/eZdravstvo/appointment/appointments-doctor/appointments.component';
import { RegisterComponent } from './components/register/register.component';
import { LoginComponent } from './components/login/login.component';
import { CompetitionsComponent } from './components/competitions/competitions/competitions.component';
import { CompetitionListComponent } from './components/competitions/competition-list/competition-list.component';
import { CompetitionAddComponent } from './components/competitions/competition-add/competition-add.component';
import { CompetitionItemComponent } from './components/competitions/competition-item/competition-item.component';
import { CompetitionViewComponent } from './components/competitions/competition-view/competition-view.component';
import { RegisterAprComponent } from './components/register-apr/register-apr.component';
import { MyAprsComponent } from './components/my-aprs/my-aprs.component';
import { AprListItemComponent } from './components/apr-list-item/apr-list-item.component';
import { WelcomeComponent } from './components/welcome/welcome.component';
import { AppointmentsRegularComponent } from './components/eZdravstvo/appointment/appointments-regular/appointments-regular.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatCardModule } from '@angular/material/card';
import { MatSelectModule } from '@angular/material/select';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon'
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatNativeDateModule} from '@angular/material/core';
import { VaccinationAddComponent } from './components/eZdravstvo/vaccination/vaccination-add/vaccination-add.component';
import { VaccinationItemComponent } from './components/eZdravstvo/vaccination/vaccination-item/vaccination-item.component';
import { VaccinationListComponent } from './components/eZdravstvo/vaccination/vaccination-list/vaccination-list.component';
import { VaccinationViewComponent } from './components/eZdravstvo/vaccination/vaccination-view/vaccination-view.component';
import { VaccinationsDoctorComponent } from './components/eZdravstvo/vaccination/vaccinations-doctor/vaccinations-doctor.component';
import { VaccinationsRegularComponent } from './components/eZdravstvo/vaccination/vaccinations-regular/vaccinations-regular.component'; 
import { MarriageComponent } from './components/marriage/marriage.component';
import { ChooseServiceComponent } from './components/choose-service/choose-service.component';
import { RegularOrAdminComponent } from './components/regular-or-admin/regular-or-admin.component'; 
import {MatSnackBarModule} from "@angular/material/snack-bar";
import { VrticAddComponent } from './components/preschool/vrtic-add/vrtic-add.component';
import { VrticListComponent } from './components/preschool/vrtic-list/vrtic-list.component';
import { VrticViewComponent } from './components/preschool/vrtic-view/vrtic-view.component';
import { VrticPocetnaComponent } from './components/preschool/vrtic-pocetna/vrtic-pocetna.component';
import { VrticItemComponent } from './components/preschool/vrtic-item/vrtic-item.component';
import { PrijavaComponent } from './components/preschool/prijava/prijava.component';
import { RegisterCrosoComponent } from './components/register-croso/register-croso.component';
import { MyCrososComponent } from './components/my-crosos/my-crosos.component';
import { CrosoListItemComponent } from './components/croso-list-item/croso-list-item.component';
import { RegisterEmployeeComponent } from './components/register-employee/register-employee.component';
import { CompanyEmployeesComponent } from './components/company-employees/company-employees.component';
import { EmployeeListItemComponent } from './components/employee-list-item/employee-list-item.component';
import { PrijavaListComponent } from './components/preschool/prijava-list/prijava-list.component';
import { PrijavaItemComponent } from './components/preschool/prijava-item/prijava-item.component';
import { PrijavaMainComponent } from './components/preschool/prijava-main/prijava-main.component';
import { AddPersonRegistryComponent } from './components/eZdravstvo/add-person-registry/add-person-registry.component'; 
import { ViewMyRegistryComponent } from './components/view-my-registry/view-my-registry.component';
import { ZdravstvenoStanjeItemComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstveno-stanje-item/zdravstveno-stanje-item.component';
import { ZdravstvenoStanjeListComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstveno-stanje-list/zdravstveno-stanje-list.component';
import { ZdravstvenoStanjeAddComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstveno-stanje-add/zdravstveno-stanje-add.component';
import { ZdravstvenaStanjaDoctorComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstvena-stanja-doctor/zdravstvena-stanja-doctor.component';
import { ZdravstvenoStanjeViewMyComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstveno-stanje-view-my/zdravstveno-stanje-view-my.component';
import { VaccinationsMyRegularComponent } from './components/eZdravstvo/vaccination/vaccinations-my-regular/vaccinations-my-regular.component';
import { UserDiedComponent } from './components/eZdravstvo/user-died/user-died.component';



@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    AppointmentAddComponent,
    AppointmentItemComponent,
    AppointmentListComponent,
    AppointmentViewComponent,
    AppointmentsDoctorComponent,
    RegisterComponent,
    LoginComponent,
    CompetitionsComponent,
    CompetitionListComponent,
    CompetitionAddComponent,
    CompetitionItemComponent,
    CompetitionViewComponent,
    RegisterAprComponent,
    MyAprsComponent,
    AprListItemComponent,
    WelcomeComponent,
    AppointmentsRegularComponent,
    VaccinationAddComponent,
    VaccinationItemComponent,
    VaccinationListComponent,
    VaccinationViewComponent,
    VaccinationsDoctorComponent,
    VaccinationsRegularComponent,
    MarriageComponent,
    ChooseServiceComponent,
    RegularOrAdminComponent,
    VrticAddComponent,
    VrticListComponent,
    VrticViewComponent,
    VrticPocetnaComponent,
    VrticItemComponent,
    PrijavaComponent,
    PrijavaListComponent,
    PrijavaItemComponent,
    PrijavaMainComponent,
    AddPersonRegistryComponent,
    RegisterCrosoComponent,
    MyCrososComponent,
    CrosoListItemComponent,
    RegisterEmployeeComponent,
    CompanyEmployeesComponent,
    EmployeeListItemComponent,
    ViewMyRegistryComponent,
    ZdravstvenoStanjeItemComponent,
    ZdravstvenoStanjeListComponent,
    ZdravstvenoStanjeAddComponent,
    ZdravstvenaStanjaDoctorComponent,
    ZdravstvenoStanjeViewMyComponent,
    VaccinationsMyRegularComponent,
    UserDiedComponent
  ],
  imports: [
    BrowserModule,
    MatToolbarModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    MatCardModule,
    MatSelectModule,
    MatButtonModule,
    MatIconModule,
    MatDatepickerModule,
    MatNativeDateModule,
    ReactiveFormsModule,
    BrowserAnimationsModule,
    MatSnackBarModule,
  ],
  providers: [{
    provide: HTTP_INTERCEPTORS,
    useClass: AuthInterceptor,
    multi: true
  }],
  bootstrap: [AppComponent]
})
export class AppModule { }
