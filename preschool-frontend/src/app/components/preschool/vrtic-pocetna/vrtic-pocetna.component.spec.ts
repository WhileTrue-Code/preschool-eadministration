import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VrticPocetnaComponent } from './vrtic-pocetna.component';

describe('VrticPocetnaComponent', () => {
  let component: VrticPocetnaComponent;
  let fixture: ComponentFixture<VrticPocetnaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VrticPocetnaComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VrticPocetnaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
