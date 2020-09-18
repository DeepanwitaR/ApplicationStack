import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { Left2Component } from './left2.component';

describe('Left2Component', () => {
  let component: Left2Component;
  let fixture: ComponentFixture<Left2Component>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ Left2Component ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(Left2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
