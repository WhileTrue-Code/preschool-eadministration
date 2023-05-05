import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AprListItemComponent } from './apr-list-item.component';

describe('AprListItemComponent', () => {
  let component: AprListItemComponent;
  let fixture: ComponentFixture<AprListItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AprListItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AprListItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
