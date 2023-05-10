import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AppointmentsRegularComponent } from './appointments-regular.component';

describe('AppointmentsRegularComponent', () => {
  let component: AppointmentsRegularComponent;
  let fixture: ComponentFixture<AppointmentsRegularComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AppointmentsRegularComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AppointmentsRegularComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
