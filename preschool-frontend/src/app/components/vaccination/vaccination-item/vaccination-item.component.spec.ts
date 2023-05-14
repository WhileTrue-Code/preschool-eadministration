import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VaccinationItemComponent } from './vaccination-item.component';

describe('VaccinationItemComponent', () => {
  let component: VaccinationItemComponent;
  let fixture: ComponentFixture<VaccinationItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VaccinationItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VaccinationItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
