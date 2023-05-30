import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VrticAddComponent } from './vrtic-add.component';

describe('VrticAddComponent', () => {
  let component: VrticAddComponent;
  let fixture: ComponentFixture<VrticAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VrticAddComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VrticAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
