import { TestBed } from '@angular/core/testing';

import { SessionStopService } from './session-stop';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';
import { StopStatus } from '../generated/models/stop-status';
import { required } from '@angular/forms/signals';

describe('SessionStop', () => {
  let service: SessionStopService;
  let httpMock:HttpTestingController; 

  beforeEach(() => {
    TestBed.configureTestingModule({

      providers:[
        SessionStopService,
        provideHttpClient(),
        provideHttpClientTesting()
      ]
    });
    service = TestBed.inject(SessionStopService);
    httpMock = TestBed.inject(HttpTestingController); 
  });

  afterEach(() => {
    httpMock.verify(); 
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('Should send POST /logs/sessions/{id}/stop success',() => {
      
      const expected:StopStatus = {success:true}; 
      service.stopService("xfg").subscribe((res) => {
        expect(res).toEqual(expected); 
        expect(res.success).toBe(true); 
      });

      const request = httpMock.expectOne('http://192.168.4.1:8080/logs/sessions/xfg/stop');
      expect(request.request.method).toBe('POST'); 
      expect(request.request.body).toEqual({}); 
      request.flush(expected); 
    }); 

  it('Should send POST /logs/sessions/{id}/stop sucess false', () =>{

      
  }); 

});
