import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { HomePage } from './home-page';
import {MockStatusService} from '../status/status.spec';
import {MockLogService} from '../logs/logs.spec'; 
import { Status } from '../../generated/model/status';
import { Log } from '../../generated/model/log';
import { StatusService } from '../../services/status.service';
import { LogService } from '../../services/log.service';

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

  it('component level test status and logs', ()=> {

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
});
