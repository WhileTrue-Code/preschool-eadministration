import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterAprComponent } from './register-apr.component';

describe('RegisterAprComponent', () => {
  let component: RegisterAprComponent;
  let fixture: ComponentFixture<RegisterAprComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisterAprComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RegisterAprComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
