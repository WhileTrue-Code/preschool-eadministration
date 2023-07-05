import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChangeCompanyDetailsComponent } from './change-company-details.component';

describe('ChangeCompanyDetailsComponent', () => {
  let component: ChangeCompanyDetailsComponent;
  let fixture: ComponentFixture<ChangeCompanyDetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChangeCompanyDetailsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ChangeCompanyDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
