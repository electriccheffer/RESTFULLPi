import { ComponentFixture, TestBed, tick } from '@angular/core/testing';
import {LogService} from '../../services/log.service';
import { Logs } from './logs';
import { Observable,of } from 'rxjs';
import {Log} from '../../generated/models/log';
import { ChangeDetectorRef, Injectable } from '@angular/core';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';
import { MockSessionStartService } from '../session-start/session-start.spec';
import { SessionStartService } from '../../services/session-start.service';

@Injectable({

  providedIn:'root',

})
export class MockLogService extends LogService{

  private logs:Log[] = [];

  setLogs(newLogs:Log[]):void{

    this.logs = newLogs;
  }

  override getLogs():Observable<Log[]>{

    return of(this.logs); 
  }

}

describe('Logs', () => {
  let component: Logs;
  let fixture: ComponentFixture<Logs>;
  let mockLogService: MockLogService ; 

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Logs],
      providers:[{

        provide:LogService,
        useClass:MockLogService

      }]
    }).compileComponents();
    mockLogService = TestBed.inject(LogService)as unknown as MockLogService;
    fixture = TestBed.createComponent(Logs);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  
  it('empty list behavior', () => {
    
    const compiled = fixture.nativeElement as HTMLElement; 
    const emptyState = compiled.querySelector('.empty-log-list');
    expect(emptyState).toBeTruthy();
    expect(emptyState?.textContent).toBe('There are no current logs');

  }); 

  it('list with one value', () => {
    
    const log: Log = { name: 'firstLog' };
    mockLogService.setLogs([log]);
    fixture = TestBed.createComponent(Logs); 
    component = fixture.componentInstance;
    fixture.detectChanges(); 
    component.ngOnInit(); 
    const compiled = fixture.nativeElement as HTMLElement;
    const cell = compiled.querySelector(`[data-testid="logName-${log.name}"]`);

    expect(cell).not.toBeNull();
    expect(cell?.textContent?.trim()).toBe('firstLog');

  }); 

  it('list with multiple values',() => {

    const log: Log = { name: 'firstLog' };
    const logTwo: Log = {name:'secondLog'}
    mockLogService.setLogs([log,logTwo]);
    fixture = TestBed.createComponent(Logs); 
    component = fixture.componentInstance;
    fixture.detectChanges(); 
    component.ngOnInit(); 
    const compiled = fixture.nativeElement as HTMLElement;
    const cell = compiled.querySelector(`[data-testid="logName-${logTwo.name}"]`);

    expect(cell).not.toBeNull();
    expect(cell?.textContent?.trim()).toBe('secondLog');
    const cellTwo = compiled.querySelector(`[data-testid="logName-${log.name}"]`);
    expect(cellTwo).not.toBeNull();
    expect(cellTwo?.textContent?.trim()).toBe('firstLog');
  })

});

describe("Logs Component LogService Integration test",() => {

  let component:Logs;
  let logService:LogService; 
  let fixture:ComponentFixture<Logs>; 
  let httpTesting: HttpTestingController; 

  beforeEach(async() => {

    await TestBed.configureTestingModule({
      imports:[Logs],
      providers:[
        LogService,
        provideHttpClient(),
        provideHttpClientTesting()
      ],

    }).compileComponents();

    httpTesting = TestBed.inject(HttpTestingController);
    fixture = TestBed.createComponent(Logs);
    component = fixture.componentInstance;
  }); 

  afterEach(() => {

    httpTesting.verify();

  });

  it('Integration Test for displaying empty logs',()=>{

    const apiResponse:Log[] = []; 

    fixture.detectChanges();
    
    const request = httpTesting.expectOne('http://192.168.4.1:8080/logs');
    expect(request.request.method).toBe("GET");
    request.flush(apiResponse);

    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement; 
    const rows = compiled.querySelector('.empty-log-list');
    expect(rows).not.toBeNull();
    expect(rows?.textContent).toBe('There are no current logs'); 
  });
  
  it('Integration Test for displaying one log',async()=>{
    const log:Log = {
      name:'logOne',
    }
    const apiResponse:Log[] = [log]; 

    fixture.detectChanges();

    const request = httpTesting.expectOne('http://192.168.4.1:8080/logs');
    expect(request.request.method).toBe("GET");
    request.flush(apiResponse);
    
    fixture.debugElement.injector.get(ChangeDetectorRef).detectChanges();
    fixture.detectChanges();
    
    const compiled = fixture.nativeElement as HTMLElement; 
    const rows = compiled.querySelector(`[data-testid="logName-${log.name}"]`);
    expect(rows).not.toBeNull();
    expect(rows?.textContent?.trim()).toBe(log.name);
  });

  it('Integration Test for displaying two logs',()=>{

    const log:Log = {
      name:'logOne',
    }

    const logTwo:Log = {

      name:'logTwo',

    }
    const apiResponse:Log[] = [log,logTwo]; 

    fixture.detectChanges();

    const request = httpTesting.expectOne('http://192.168.4.1:8080/logs');
    expect(request.request.method).toBe("GET");
    request.flush(apiResponse);
    
    fixture.debugElement.injector.get(ChangeDetectorRef).detectChanges();
    fixture.detectChanges();
    
    const compiled = fixture.nativeElement as HTMLElement; 
    const rows = compiled.querySelector(`[data-testid="logName-${log.name}"]`);
    expect(rows).not.toBeNull();
    expect(rows?.textContent?.trim()).toBe(log.name);

    const rowsTwo = compiled.querySelector(`[data-testid="logName-${logTwo.name}"]`);
    expect(rowsTwo).not.toBeNull();
    expect(rowsTwo?.textContent?.trim()).toBe(logTwo.name);
  });

  

}); 

describe("Logs Session Service with Mock SesssionStartService",() =>{

  let component: Logs;
  let fixture: ComponentFixture<Logs>;
  let mockLogService: MockLogService ; 
  let mockSessionStartService: MockSessionStartService; 

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Logs],
      providers:[{

        provide:LogService,
        useClass:MockLogService,
        },{
        provide:SessionStartService,
        useClass:MockSessionStartService
        }
      ]
    }).compileComponents();
    mockLogService = TestBed.inject(LogService)as unknown as MockLogService;
    mockSessionStartService = TestBed.inject(SessionStartService) as unknown as MockSessionStartService; 
    fixture = TestBed.createComponent(Logs);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });


  it('Integration Test for LogsComponent with SessionStartService',() => {

    const expectedLog:Log = {name:'09_01_2026.gpx'}; 
    const expectedLogs:Log[] = [expectedLog]; 
    fixture.detectChanges();

    mockLogService.setLogs(expectedLogs); 
    mockSessionStartService.startSession().subscribe();
    fixture.detectChanges(); 
    const compiled = fixture.nativeElement as HTMLElement; 
    const rows = compiled.querySelector(`[data-testid="logName-${expectedLog.name}"]`);
    expect(rows).not.toBeNull();
    expect(rows?.textContent?.trim()).toBe(expectedLog.name); 

  });

});

describe("Integration Test with SessionStartService",()=>{

  let component:Logs; 
  let fixture:ComponentFixture<Logs>; 
  let sessionStartService:SessionStartService; 
  let mockLogService:MockLogService; 
  let httpTesting: HttpTestingController; 

  beforeEach(async() => {

    await TestBed.configureTestingModule({
      imports:[Logs],
      providers:[
        provideHttpClient(),
        provideHttpClientTesting(),
        SessionStartService,
        {provide: LogService,useClass:MockLogService}
      ]
    }).compileComponents();

    sessionStartService = TestBed.inject(SessionStartService);
    httpTesting = TestBed.inject(HttpTestingController);
    mockLogService = TestBed.inject(MockLogService);
    fixture = TestBed.createComponent(Logs);
    component = fixture.componentInstance; 
  });
    afterEach(() => {

      httpTesting.verify();

    }); 

    it("Should Initialize and Create",() => {

      expect(component).toBeTruthy(); 

    });

  }); 

