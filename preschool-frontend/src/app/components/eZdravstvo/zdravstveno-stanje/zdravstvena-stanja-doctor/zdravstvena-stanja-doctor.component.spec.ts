import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ZdravstvenaStanjaDoctorComponent } from './zdravstvena-stanja-doctor.component';

describe('ZdravstvenaStanjaDoctorComponent', () => {
  let component: ZdravstvenaStanjaDoctorComponent;
  let fixture: ComponentFixture<ZdravstvenaStanjaDoctorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ZdravstvenaStanjaDoctorComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ZdravstvenaStanjaDoctorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
