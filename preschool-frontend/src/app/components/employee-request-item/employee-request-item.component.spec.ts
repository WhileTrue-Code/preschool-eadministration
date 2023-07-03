import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EmployeeRequestItemComponent } from './employee-request-item.component';

describe('EmployeeRequestItemComponent', () => {
  let component: EmployeeRequestItemComponent;
  let fixture: ComponentFixture<EmployeeRequestItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EmployeeRequestItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EmployeeRequestItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
