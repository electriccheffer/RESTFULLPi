import { ComponentFixture, TestBed } from '@angular/core/testing';
import {SessionStartService} from '../../services/session-start.service';
import { SessionStart } from './session-start';
import { Observable, of } from 'rxjs';
import {StartStatus} from '../../generated/models/start-status'

export class MockSessionStartService extends SessionStartService{

  status:StartStatus = {name:''};

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
  let mockSessionStartService:MockSessionStartService = new MockSessionStartService();

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SessionStart],
      providers: [

        {provide: SessionStartService, useValue:mockSessionStartService}

      ]
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
    
    const expectedFilename = `gps_log_${mm}-${dd}-${yy}_${hh}-${min}-${ss}.log`;

    const startStatus:StartStatus = {name:expectedFilename};
    mockSessionStartService.setSession(startStatus);
    
    fixture.detectChanges(); 
    
    // check for button 
    const compiled = fixture.nativeElement as HTMLElement; 
    const button = compiled.querySelector('button');
    const emptyName = compiled.querySelector('p.file-name-display'); 
    expect(emptyName).toBeNull(); 
    expect(button).not.toBeNull();
    
    // press button check for file name 
    button?.click(); 

    fixture.detectChanges(); 
    const fileName = compiled.querySelector('p.file-name-display');
    expect(fileName).not.toBeNull();
    expect(fileName?.textContent?.trim()).toContain(expectedFilename);   
  });
  
  
});
