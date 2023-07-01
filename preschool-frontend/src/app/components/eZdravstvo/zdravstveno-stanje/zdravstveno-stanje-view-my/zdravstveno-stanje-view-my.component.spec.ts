import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ZdravstvenoStanjeViewMyComponent } from './zdravstveno-stanje-view-my.component';

describe('ZdravstvenoStanjeViewMyComponent', () => {
  let component: ZdravstvenoStanjeViewMyComponent;
  let fixture: ComponentFixture<ZdravstvenoStanjeViewMyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ZdravstvenoStanjeViewMyComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ZdravstvenoStanjeViewMyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
