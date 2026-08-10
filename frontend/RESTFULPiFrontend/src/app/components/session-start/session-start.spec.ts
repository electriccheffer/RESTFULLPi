import { ComponentFixture, TestBed } from '@angular/core/testing';
import {SessionStartService} from '../../services/session-start.service';
import { SessionStart } from './session-start';
import { Observable, of } from 'rxjs';
import {StartStatus} from '../../generated/models/start-status'

class MockSessionStartService extends SessionStartService{

  status:StartStatus = {name:'1234'}

  override startSession(): Observable<StartStatus> {
    return of(this.status);
  }

}

describe('SessionStart', () => {
  let component: SessionStart;
  let fixture: ComponentFixture<SessionStart>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SessionStart],
    }).compileComponents();

    fixture = TestBed.createComponent(SessionStart);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });


});
