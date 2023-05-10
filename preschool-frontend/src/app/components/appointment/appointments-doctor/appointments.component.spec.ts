import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AppointmentsDoctorComponent } from './appointments.component';

describe('AppointmentsComponent', () => {
  let component: AppointmentsDoctorComponent;
  let fixture: ComponentFixture<AppointmentsDoctorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AppointmentsDoctorComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AppointmentsDoctorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
