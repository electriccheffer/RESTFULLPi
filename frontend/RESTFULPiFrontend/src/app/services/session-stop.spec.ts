import { TestBed } from '@angular/core/testing';

import { SessionStopService } from './session-stop';

describe('SessionStop', () => {
  let service: SessionStopService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SessionStopService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
