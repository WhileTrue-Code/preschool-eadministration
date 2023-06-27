import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ZdravstvenoStanjeAddComponent } from './zdravstveno-stanje-add.component';

describe('ZdravstvenoStanjeAddComponent', () => {
  let component: ZdravstvenoStanjeAddComponent;
  let fixture: ComponentFixture<ZdravstvenoStanjeAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ZdravstvenoStanjeAddComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ZdravstvenoStanjeAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
