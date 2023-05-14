import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VaccinationsDoctorComponent } from './vaccinations-doctor.component';

describe('VaccinationsDoctorComponent', () => {
  let component: VaccinationsDoctorComponent;
  let fixture: ComponentFixture<VaccinationsDoctorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VaccinationsDoctorComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VaccinationsDoctorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
