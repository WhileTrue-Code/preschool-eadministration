import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ZdravstvenoStanjeItemComponent } from './zdravstveno-stanje-item.component';

describe('ZdravstvenoStanjeItemComponent', () => {
  let component: ZdravstvenoStanjeItemComponent;
  let fixture: ComponentFixture<ZdravstvenoStanjeItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ZdravstvenoStanjeItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ZdravstvenoStanjeItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
