import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VrticListComponent } from './vrtic-list.component';

describe('VrticListComponent', () => {
  let component: VrticListComponent;
  let fixture: ComponentFixture<VrticListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VrticListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VrticListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
