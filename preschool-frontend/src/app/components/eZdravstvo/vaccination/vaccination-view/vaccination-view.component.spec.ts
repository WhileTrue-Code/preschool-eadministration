import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VaccinationViewComponent } from './vaccination-view.component';

describe('VaccinationViewComponent', () => {
  let component: VaccinationViewComponent;
  let fixture: ComponentFixture<VaccinationViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VaccinationViewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VaccinationViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
