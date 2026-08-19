import { TestBed } from "@angular/core/testing";
import { provideHttpClient } from "@angular/common/http";
import { provideHttpClientTesting,HttpTestingController } from "@angular/common/http/testing";
import { SessionStartService } from "./session-start.service";
import { StartStatus } from "../generated/models";

describe('SessionStartService',() => {

    let service:SessionStartService;
    let httpMock:HttpTestingController; 
    
    beforeEach(() => {

        TestBed.configureTestingModule({
            providers:[
                SessionStartService,
                provideHttpClient(),
                provideHttpClientTesting()
            ]
        });
    
        service = TestBed.inject(SessionStartService); 
        httpMock = TestBed.inject(HttpTestingController);


    });

    afterEach(()=>{
        httpMock.verify(); 
    }); 
    
    it('Should create',()=>{

        expect(service).toBeTruthy(); 

    }); 

    it('Should send POST /logs/sessions request and return StartStatus object',() => {

        const expected:StartStatus = {name:'gps_log_08_19_2026.log'};
        
        service.startSession().subscribe((res) => {
            expect(res).toEqual(expected); 
            expect(res.name).toBe(expected.name); 
        }); 

        const request = httpMock.expectOne('http://192.168.4.1:8080/logs/sessions');
        expect(request.request.method).toBe('POST');
        expect(request.request.body).toEqual({});
        request.flush(expected); 
    }); 

});

