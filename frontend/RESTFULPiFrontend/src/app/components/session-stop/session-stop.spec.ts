import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SessionStop } from './session-stop';
import { SessionStopService } from '../../services/session-stop';
import { Observable,of } from 'rxjs';
import { StopStatus } from '../../generated/models/stop-status';

class MockSessionStopService extends SessionStopService{
  private status:StopStatus = {success:false}; 

  override stopService():Observable<StopStatus>{
    return of(this.status)
  }
  setStopStatus(status:StopStatus){
    this.status = status; 
  }
}

describe('SessionStop', () => {
  let component: SessionStop;
  let fixture: ComponentFixture<SessionStop>;
  let service: SessionStopService; 
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SessionStop],
      providers:[MockSessionStopService,
                {provide:SessionStopService,
                  useExisting:MockSessionStopService}
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(SessionStop);
    component = fixture.componentInstance;
    service = TestBed.inject(SessionStopService);
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
