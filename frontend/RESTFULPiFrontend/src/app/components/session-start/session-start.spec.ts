import { ComponentFixture, TestBed } from '@angular/core/testing';
import {SessionStartService} from '../../services/session-start.service';
import { SessionStart } from './session-start';
import { Observable, of } from 'rxjs';
import {StartStatus} from '../../generated/models/start-status'

class MockSessionStartService extends SessionStartService{

  status:StartStatus = {name:'1234'}

  setSession(status:StartStatus){

    this.status = status; 

  }

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

  it('file created success',()=>{

    const now = new Date();

    const pad = (n: number) => String(n).padStart(2, '0');
    
    const mm = pad(now.getMonth() + 1);
    const dd = pad(now.getDate());
    const yy = String(now.getFullYear()).slice(-2);
    
    const hh = pad(now.getHours());
    const min = pad(now.getMinutes());
    const ss = pad(now.getSeconds());
    
    const filename = `gps_log_${mm}-${dd}-${yy}_${hh}-${min}-${ss}.log`;

    const startSessionStatus:MockSessionStartService = new MockSessionStartService();
    const startStatus:StartStatus = {name:filename};
    startSessionStatus.setSession(startStatus);
    
    fixture.detectChanges(); 
    


  });
  
  
});
