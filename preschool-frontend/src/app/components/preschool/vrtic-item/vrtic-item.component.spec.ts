import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VrticItemComponent } from './vrtic-item.component';

describe('VrticItemComponent', () => {
  let component: VrticItemComponent;
  let fixture: ComponentFixture<VrticItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VrticItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VrticItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
