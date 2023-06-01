import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddPersonRegistryComponent } from './add-person-registry.component';

describe('AddPersonRegistryComponent', () => {
  let component: AddPersonRegistryComponent;
  let fixture: ComponentFixture<AddPersonRegistryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddPersonRegistryComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddPersonRegistryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
