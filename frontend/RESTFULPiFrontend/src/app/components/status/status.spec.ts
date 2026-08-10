import { ComponentFixture, TestBed } from '@angular/core/testing';
import {StatusService} from '../../services/status.service';
import { StatusComponent} from './status';
import {Status} from '../../generated/model/status';
import { Observable,of } from 'rxjs';
import { Injectable } from '@angular/core';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';

@Injectable()
export class MockStatusService extends StatusService{

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

    expect(component.device()).toBe('RESTFULPi');
    expect(component.status()).toBe('Online');    
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
    expect(component.device()).toBe('RESTFULPi');
    expect(component.status()).toBe('Offline');
  })
});

describe('StatusIntegrationTest',() =>{
  let component:StatusComponent;
  let fixture:ComponentFixture<StatusComponent>;
  let httpTesting: HttpTestingController; 

  beforeEach(async() => {

    await TestBed.configureTestingModule({

      imports:[StatusComponent],
      providers:[
        StatusService,
        provideHttpClient(),
        provideHttpClientTesting(),
      ],

    }).compileComponents();

    fixture = TestBed.createComponent(StatusComponent);
    component = fixture.componentInstance;
    httpTesting = TestBed.inject(HttpTestingController);

  });

  afterEach(() => {

    httpTesting.verify();

  });

  it('Integration Test Displays online through StatusServer',()=>{

    fixture.detectChanges(); 
    const request = httpTesting.expectOne('http://192.168.4.1:8080/status');
    expect(request.request.method).toBe('GET');
    const apiResponse : Status = {
        device:'RESTFULPi',
        status:'Online'

    }; 

    request.flush(apiResponse);
    fixture.detectChanges();

    expect(component.device()).toBe('RESTFULPi');
    expect(component.status()).toBe('Online');

  });

  it('Integration Test Shows offline when server not reachable',()=>{

    fixture.detectChanges();
    const request = httpTesting.expectOne('http://192.168.4.1:8080/status');
    expect(request.request.method).toBe('GET');
    request.error(new ProgressEvent('error'),{

      status: 0,
      statusText: 'Unknown Error', 
    }); 
    fixture.detectChanges(); 
    expect(component.device()).toBe('RESTFULPi');
    expect(component.status()).toBe('Server Unavailable, Unknown Status');
  }); 
}); 