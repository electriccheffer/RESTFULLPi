import { ComponentFixture, TestBed } from '@angular/core/testing';
import {StatusService} from '../../services/status.service';
import { StatusComponent} from './status';
import {Status} from '../../generated/model/status';
import { Observable,of } from 'rxjs';

class MockStatusService extends StatusService{

  private status: Status = {

    device: 'RESTFULPi',
    status: 'Online'
  }

  setStatus(newStatus:Status):void{

    this.status = newStatus; 
    
  }
  
  override getStatus():Observable<Status>{

    return of(this.status); 

  }

}

describe('Status', () => {
  let component: StatusComponent;
  let fixture: ComponentFixture<StatusComponent>;
  let mockStatusServer:MockStatusService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [StatusComponent],
      providers:[
        {
          provide:StatusService,
          useClass:MockStatusService
        }    
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(StatusComponent);
    component = fixture.componentInstance;
    mockStatusServer = TestBed.inject(StatusService) as MockStatusService;
    await fixture.whenStable();
  });

  
  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('successful service', () => {

    expect(component.device).toBe('RESTFULPi');
    expect(component.status).toBe('Online');    
  })

  it('non-successful service', () => {
    const nonSuccessStatus: Status = {
      device:'RESTFULPi', 
      status:'Offline'

    }
    mockStatusServer.setStatus(nonSuccessStatus);
    fixture = TestBed.createComponent(StatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    expect(component.device).toBe('RESTFULPi');
    expect(component.status).toBe('Offline');
  })
});
