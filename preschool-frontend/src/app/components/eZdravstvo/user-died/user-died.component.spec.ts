import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserDiedComponent } from './user-died.component';

describe('UserDiedComponent', () => {
  let component: UserDiedComponent;
  let fixture: ComponentFixture<UserDiedComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UserDiedComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UserDiedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
