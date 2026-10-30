import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SessionStop } from './session-stop';
import { SessionStopService } from '../../services/session-stop';
import { Observable,of } from 'rxjs';
import { StopStatus } from '../../generated/models/stop-status';
import { Injectable } from '@angular/core';
import { Log } from '../../generated/models';

@Injectable()
class MockSessionStopService extends SessionStopService{
  private status:StopStatus = {success:false}; 

  override stopService(id:string):Observable<StopStatus>{
    return of(this.status)
  }
  setStopStatus(status:StopStatus){
    this.status = status; 
  }
}

describe('SessionStop', () => {
  let component: SessionStop;
  let fixture: ComponentFixture<SessionStop>;
  let service: MockSessionStopService;
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
    service = TestBed.inject(MockSessionStopService);
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('Session stop success',() => {

    // get the button component
    component.log = {
      id:"123",
      name:"testfile",
      active:true

    }
    fixture.detectChanges();
    const stopStatus = {success:true}; 
    service.setStopStatus(stopStatus); 
    const compiled = fixture.nativeElement as HTMLElement;
    const button = compiled.querySelector('button'); 
    expect(button).not.toBeNull(); 
    button?.click(); 
    fixture.detectChanges();
    const status = compiled.querySelector('.stop-status p'); 
    expect(status)?.toBeTruthy(); 
    expect(status?.textContent)?.toContain('Session Stopped!'); 
  }); 

  it('Session stop failure false',() => {
    
  }); 

  it('Sessions stop error',() => {

  }); 

});
