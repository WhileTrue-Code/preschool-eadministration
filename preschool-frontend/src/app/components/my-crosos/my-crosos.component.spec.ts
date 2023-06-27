import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyCrososComponent } from './my-crosos.component';

describe('MyCrososComponent', () => {
  let component: MyCrososComponent;
  let fixture: ComponentFixture<MyCrososComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MyCrososComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MyCrososComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
