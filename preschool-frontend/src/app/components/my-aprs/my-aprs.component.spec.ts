import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyAprsComponent } from './my-aprs.component';

describe('MyAprsComponent', () => {
  let component: MyAprsComponent;
  let fixture: ComponentFixture<MyAprsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MyAprsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MyAprsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
