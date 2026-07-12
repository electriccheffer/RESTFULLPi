import { ComponentFixture, TestBed } from '@angular/core/testing';
import {StatusService} from '../../services/status.service';

import { StatusComponent} from './status';

class MockStatusService extends StatusService{

  status = {

    device: 'RESTFULPi',
    status: 'Online'
  }

  getStatus(){

    return {

      device:this.status.device, 
      status:this.status.status

    } 

  }

}

describe('Status', () => {
  let component: StatusComponent;
  let fixture: ComponentFixture<StatusComponent>;
  let mockStatusServer:StatusService;

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
    mockStatusServer = TestBed.inject(StatusService);
    await fixture.whenStable();
  });

  
  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('successful service', () => {

    expect(component.device).toBe('RESTFULPi');
    expect(component.status).toBe('Online');    
  })

});
