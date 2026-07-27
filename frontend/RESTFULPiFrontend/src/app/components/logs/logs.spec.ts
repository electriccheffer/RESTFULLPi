import { ComponentFixture, TestBed } from '@angular/core/testing';
import {LogService} from '../../services/log.service';
import { Logs } from './logs';
import { Observable,of } from 'rxjs';
import {Log} from '../../generated/model/log';
import { Injectable } from '@angular/core';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';
@Injectable({

  providedIn:'root',

})
class MockLogService extends LogService{

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
    
    const log: Log = { name: 'firstLog', date: '2026-07-26T11:18:29Z' };
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

    const log: Log = { name: 'firstLog', date: '2026-07-26T11:18:29Z' };
    const logTwo: Log = {name:'secondLog', date:'2026-07-26T11:18:29Z'}
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

    fixture = TestBed.createComponent(Logs); 
    component = fixture.componentInstance;
    httpTesting = TestBed.inject(HttpTestingController);

  }); 

  afterEach(() => {

    httpTesting.verify();

  });

  it('Integration Test for displaying empty logs',()=>{

    const apiResponse:Log[] = []; 

    fixture.detectChanges(); 
    
    const request = httpTesting.expectOne('http://192.168.4.1/logs');
    expect(request.request.method).toBe("GET");
    request.flush(apiResponse);

    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement; 
    const rows = compiled.querySelector('.empty-log-list');
    expect(rows).not.toBeNull();
    expect(rows?.textContent).toBe('There are no current logs'); 
  });
  
  it('Integration Test for displaying one log',()=>{



  });

  it('Integration Test for displaying two logs',()=>{



  });

}); 
