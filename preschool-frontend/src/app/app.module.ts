import { NgModule } from '@angular/core';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatCardModule } from '@angular/material/card';
import { MatNativeDateModule } from '@angular/material/core';
import { MatButtonModule} from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import { MatDividerModule } from '@angular/material/divider';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatCheckboxModule } from '@angular/material/checkbox';
import {MatDatepickerModule} from '@angular/material/datepicker';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { AuthInterceptor } from './services/auth.interceptor';
import { AppointmentAddComponent } from './components/appointment/appointment-add/appointment-add.component';
import { AppointmentItemComponent } from './components/appointment/appointment-item/appointment-item.component';
import { AppointmentListComponent } from './components/appointment/appointment-list/appointment-list.component';
import { AppointmentViewComponent } from './components/appointment/appointment-view/appointment-view.component';
import { AppointmentsComponent } from './components/appointment/appointments/appointments.component';
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

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    AppointmentAddComponent,
    AppointmentItemComponent,
    AppointmentListComponent,
    AppointmentViewComponent,
    AppointmentsComponent,
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
    WelcomeComponent
  ],
  imports: [
    BrowserModule,
    MatToolbarModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatToolbarModule,
    MatIconModule,
    MatCardModule,
    MatFormFieldModule,
    MatNativeDateModule,
    MatDatepickerModule,
    MatSelectModule,
    MatDividerModule,
    BrowserAnimationsModule,
    MatCheckboxModule,
    MatDatepickerModule,
    MatNativeDateModule,
  ],
  providers: [{
    provide: HTTP_INTERCEPTORS,
    useClass: AuthInterceptor,
    multi: true
  }],
  bootstrap: [AppComponent]
})
export class AppModule { }
