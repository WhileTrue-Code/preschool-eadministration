import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterCrosoComponent } from './register-croso.component';

describe('RegisterCrosoComponent', () => {
  let component: RegisterCrosoComponent;
  let fixture: ComponentFixture<RegisterCrosoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisterCrosoComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RegisterCrosoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
