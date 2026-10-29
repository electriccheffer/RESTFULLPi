import { TestBed } from '@angular/core/testing';

import { SessionStop } from './session-stop';

describe('SessionStop', () => {
  let service: SessionStop;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SessionStop);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
