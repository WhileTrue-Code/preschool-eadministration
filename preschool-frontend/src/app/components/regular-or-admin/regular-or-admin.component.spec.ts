import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegularOrAdminComponent } from './regular-or-admin.component';

describe('RegularOrAdminComponent', () => {
  let component: RegularOrAdminComponent;
  let fixture: ComponentFixture<RegularOrAdminComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegularOrAdminComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RegularOrAdminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
