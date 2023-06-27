import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VaccinationsRegularComponent } from './vaccinations-regular.component';

describe('VaccinationsRegularComponent', () => {
  let component: VaccinationsRegularComponent;
  let fixture: ComponentFixture<VaccinationsRegularComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VaccinationsRegularComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VaccinationsRegularComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
