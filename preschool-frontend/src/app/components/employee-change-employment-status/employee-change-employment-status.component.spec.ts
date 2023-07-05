import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EmployeeChangeEmploymentStatusComponent } from './employee-change-employment-status.component';

describe('EmployeeChangeEmploymentStatusComponent', () => {
  let component: EmployeeChangeEmploymentStatusComponent;
  let fixture: ComponentFixture<EmployeeChangeEmploymentStatusComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EmployeeChangeEmploymentStatusComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EmployeeChangeEmploymentStatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
