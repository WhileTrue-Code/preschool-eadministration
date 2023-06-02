import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PrijavaMainComponent } from './prijava-main.component';

describe('PrijavaMainComponent', () => {
  let component: PrijavaMainComponent;
  let fixture: ComponentFixture<PrijavaMainComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PrijavaMainComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PrijavaMainComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
