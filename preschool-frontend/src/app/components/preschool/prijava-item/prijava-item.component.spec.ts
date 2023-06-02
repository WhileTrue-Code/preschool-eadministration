import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PrijavaItemComponent } from './prijava-item.component';

describe('PrijavaItemComponent', () => {
  let component: PrijavaItemComponent;
  let fixture: ComponentFixture<PrijavaItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PrijavaItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PrijavaItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
