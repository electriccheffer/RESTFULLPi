import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SessionStop } from './session-stop';

describe('SessionStop', () => {
  let component: SessionStop;
  let fixture: ComponentFixture<SessionStop>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SessionStop],
    }).compileComponents();

    fixture = TestBed.createComponent(SessionStop);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
