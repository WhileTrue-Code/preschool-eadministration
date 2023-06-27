import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CrosoListItemComponent } from './croso-list-item.component';

describe('CrosoListItemComponent', () => {
  let component: CrosoListItemComponent;
  let fixture: ComponentFixture<CrosoListItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CrosoListItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CrosoListItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
