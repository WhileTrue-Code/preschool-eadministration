import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HealthStatusAddComponent } from './health-status-add.component';

describe('HealthStatusAddComponent', () => {
  let component: HealthStatusAddComponent;
  let fixture: ComponentFixture<HealthStatusAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HealthStatusAddComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HealthStatusAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
