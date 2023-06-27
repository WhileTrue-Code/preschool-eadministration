import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewMyRegistryComponent } from './view-my-registry.component';

describe('ViewMyRegistryComponent', () => {
  let component: ViewMyRegistryComponent;
  let fixture: ComponentFixture<ViewMyRegistryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewMyRegistryComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewMyRegistryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
