import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PrijavaListComponent } from './prijava-list.component';

describe('PrijavaListComponent', () => {
  let component: PrijavaListComponent;
  let fixture: ComponentFixture<PrijavaListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PrijavaListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PrijavaListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
