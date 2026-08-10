import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { HomePage } from './home-page';
import {MockStatusService} from '../status/status.spec';
import {MockLogService} from '../logs/logs.spec'; 
import { Status } from '../../generated/model/status';
import { Log } from '../../generated/model/log';
import { StatusService } from '../../services/status.service';
import { LogService } from '../../services/log.service';
import { ChangeDetectorRef } from '@angular/core';

describe('HomePage component level test', () => {
  let component: HomePage;
  let fixture: ComponentFixture<HomePage>;
  let mockStatusService:MockStatusService; 
  let mockLogService:MockLogService;

  beforeEach(async () => {
   
    await TestBed.configureTestingModule({
      imports: [HomePage],
      providers:[
        provideHttpClient(),
        provideHttpClientTesting(),
        {provide:StatusService,useClass:MockStatusService},
        {provide:LogService,useClass:MockLogService}
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(HomePage);
    mockStatusService = TestBed.inject(StatusService) as MockStatusService; 
    mockLogService = TestBed.inject(LogService) as MockLogService; 
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('component level test status and no logs', ()=> {

    const status:Status = {
        device:'RESTFULPi',
        status:'Online'

    };
    const log:Log = {};
    const logs:Log[] = [log]; 
    mockStatusService.setStatus(status);
    mockLogService.setLogs(logs);
    fixture.detectChanges();

    const compiled = fixture.nativeElement as HTMLElement; 
    const statusText = compiled.querySelector('status')?.textContent;
    expect(statusText).toContain("RESTFULPi");
    expect(statusText).toContain("Online"); 
    
    const logsElement = compiled.querySelector('.empty-log-list')?.textContent;
    expect(logsElement).not.toBeNull();
    expect(logsElement).toContain('There are no current logs');

  });

  it('component level test status and  multiple logs',() => {

    const status:Status = {
      device:'RESTFULPi',
      status:'Online'

    };
    const log:Log = {name:'logOne'};
    const logTwo:Log = {name:'logTwo'};
    const logs:Log[] = [log,logTwo]; 

    fixture = TestBed.createComponent(HomePage);
    mockStatusService.setStatus(status);
    mockLogService.setLogs(logs);
    fixture.detectChanges();
    fixture.debugElement.injector.get(ChangeDetectorRef).detectChanges();
    fixture.detectChanges();
    
    const compiled = fixture.nativeElement as HTMLElement; 
    const statusText = compiled.querySelector('status')?.textContent;
    expect(statusText).toContain("RESTFULPi");
    expect(statusText).toContain("Online"); 
    
    const logsElement = compiled.querySelector(`[data-testid="logName-${log.name}"]`);
    expect(logsElement).not.toBeNull(); 
    expect(logsElement?.textContent?.trim()).toBe(log.name); 

    const logsElementTwo = compiled.querySelector(`[data-testid="logName-${logTwo.name}"]`);
    expect(logsElementTwo).not.toBeNull(); 
    expect(logsElementTwo?.textContent?.trim()).toBe(logTwo.name); 

  });

});

describe("HomePage Component Integration Test",() => {

  let component:HomePage; 
  let fixture:ComponentFixture<HomePage>;
  let httpTesting:HttpTestingController; 

  beforeEach(async() => {

    await TestBed.configureTestingModule({

      imports:[HomePage],
      providers:[
        provideHttpClient(),
        provideHttpClientTesting()
      ]

    }).compileComponents();

    httpTesting = TestBed.inject(HttpTestingController); 
    fixture = TestBed.createComponent(HomePage); 
    component =fixture.componentInstance;
  });

  afterEach(() => {

    httpTesting.verify(); 

  });

  it('Integration Test no logs for homepage component', () => {

    const logs:Log[] = []; 

    fixture.detectChanges();

    const logsRequest = httpTesting.expectOne('http://192.168.4.1:8080/logs');
    expect(logsRequest.request.method).toBe("GET"); 
    logsRequest.flush(logs);

    const statusRequest = httpTesting.expectOne('http://192.168.4.1:8080/status'); 
    expect(statusRequest.request.method).toBe("GET"); 
    statusRequest.flush({device:"RESTFULPi",status:"Online"});

    fixture.detectChanges();
    
    const compiled = fixture.nativeElement as HTMLElement;
    const statusRows = compiled.querySelector('status')?.textContent; 
    expect(statusRows).not.toBeNull(); 
    console.log(statusRows); 
    expect(statusRows).toContain("RESTFULPi");
    expect(statusRows).toContain("Online"); 
    
    const logsRows = compiled.querySelector('.empty-log-list')?.textContent; 
    expect(logsRows).not.toBeNull();
    expect(logsRows).toBe('There are no current logs');

  });

  it('Integration Test with logs for homepage component',() => {

    const log:Log = {name:"logOne"}; 
    const logTwo:Log = {name:"logTwo"}; 
    const logs:Log[] = [log,logTwo];

    fixture.detectChanges(); 
    
    const logsRequest = httpTesting.expectOne('http://192.168.4.1:8080/logs');
    expect(logsRequest.request.method).toBe("GET"); 
    logsRequest.flush(logs);

    const statusRequest = httpTesting.expectOne('http://192.168.4.1:8080/status'); 
    expect(statusRequest.request.method).toBe("GET"); 
    statusRequest.flush({device:"RESTFULPi",status:"Online"});

    fixture.detectChanges();
    
    const compiled = fixture.nativeElement as HTMLElement;
    const statusRows = compiled.querySelector('status')?.textContent; 
    expect(statusRows).not.toBeNull(); 
    console.log(statusRows); 
    expect(statusRows).toContain("RESTFULPi");
    expect(statusRows).toContain("Online"); 
    
    const logsRows = compiled.querySelector(`[data-testid="logName-${log.name}"]`)?.textContent; 
    expect(logsRows).not.toBeNull();
    expect(logsRows).toBe(log.name); 
  });
}); 
