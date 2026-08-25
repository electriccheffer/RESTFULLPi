import { ComponentFixture, TestBed } from '@angular/core/testing';
import {SessionStartService} from '../../services/session-start.service';
import { SessionStart } from './session-start';
import { Observable, of,throwError } from 'rxjs';
import {StartStatus} from '../../generated/models/start-status'
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { provideHttpClient } from '@angular/common/http';

export class MockSessionStartService extends SessionStartService{

  status:StartStatus = {name:''};
  errorMessage:string = ''; 
  
  constructor(){

    super(null!);

  }

  setSession(status:StartStatus):void{

    this.status = status; 

  }

  setError(errorMessage:string){

    this.errorMessage = errorMessage; 

  }

  override startSession(): Observable<StartStatus> {
    if(this.errorMessage !== ''){
      return throwError( () => new Error(this.errorMessage)); 
    }
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
  
  it('file creation error',() => {

    const expectedErrorMessage = 'Error starting file check the pi';
    mockSessionStartService.setError(expectedErrorMessage);
    fixture.detectChanges();

    const compiled = fixture.nativeElement as HTMLElement;
    const button = compiled.querySelector('button'); 
    expect(button).not.toBeNull();
    
    button?.click(); 

    fixture.detectChanges();
    const fileName = compiled.querySelector('p.file-name-display'); 
    expect(fileName).toBeNull();

    const errorInfo = compiled.querySelector('p.error-message');
    expect(errorInfo).not.toBeNull(); 
    expect(errorInfo?.textContent?.trim()).toBe(expectedErrorMessage);

  });


});


describe('SessionStart service integration test',() => {

  let component: SessionStart; 
  let fixture: ComponentFixture<SessionStart>; 
  let httpTesting: HttpTestingController; 

  beforeEach(async() => {

    await TestBed.configureTestingModule({

      imports:[SessionStart],
      providers: [
        SessionStartService,
        provideHttpClient(),
        provideHttpClientTesting()
      ]

    }).compileComponents(); 

    httpTesting = TestBed.inject(HttpTestingController); 
    fixture = TestBed.createComponent(SessionStart);
    component = fixture.componentInstance; 
  }); 

  afterEach(() => {

    httpTesting.verify(); 

  }); 

  it('Integration test of SessionStart ', () => {
    
    const fileName = '25_09_26.gpx';
    const apiResponse:StartStatus = {name:fileName};

    const button = fixture.nativeElement.querySelector('.start-session-button');
    expect(button).toBeTruthy();
    button.click();

    const request = httpTesting.expectOne('http://192.168.4.1:8080/logs/sessions');
    expect(request.request.method).toBe("POST"); 
    request.flush(apiResponse); 

    fixture.detectChanges(); 

    const pageData = fixture.nativeElement.querySelector('.file-name-display'); 
    expect(pageData).toBeTruthy();
    expect(pageData.textContent).toContain(fileName);

  }); 

}); 