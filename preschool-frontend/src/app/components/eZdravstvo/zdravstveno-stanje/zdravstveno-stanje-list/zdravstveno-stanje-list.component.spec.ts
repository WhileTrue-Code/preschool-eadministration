import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ZdravstvenoStanjeListComponent } from './zdravstveno-stanje-list.component';

describe('ZdravstvenoStanjeListComponent', () => {
  let component: ZdravstvenoStanjeListComponent;
  let fixture: ComponentFixture<ZdravstvenoStanjeListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ZdravstvenoStanjeListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ZdravstvenoStanjeListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
