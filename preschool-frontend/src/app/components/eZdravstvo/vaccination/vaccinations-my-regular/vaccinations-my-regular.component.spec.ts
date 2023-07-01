import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VaccinationsMyRegularComponent } from './vaccinations-my-regular.component';

describe('VaccinationsMyRegularComponent', () => {
  let component: VaccinationsMyRegularComponent;
  let fixture: ComponentFixture<VaccinationsMyRegularComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VaccinationsMyRegularComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VaccinationsMyRegularComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
