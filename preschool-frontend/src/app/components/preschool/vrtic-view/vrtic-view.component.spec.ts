import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VrticViewComponent } from './vrtic-view.component';

describe('VrticViewComponent', () => {
  let component: VrticViewComponent;
  let fixture: ComponentFixture<VrticViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VrticViewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VrticViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
