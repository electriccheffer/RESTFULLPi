import { HttpTestingController, provideHttpClientTesting } from "@angular/common/http/testing";
import { LogService } from "./log.service";
import { TestBed } from "@angular/core/testing";
import { provideHttpClient } from "@angular/common/http";
import { Log } from "../generated/model/log";

describe('LogService',()=>{

    let logService:LogService;
    let httpTesting:HttpTestingController;

    beforeEach(()=>{

        TestBed.configureTestingModule({

            providers:[
                provideHttpClient(),
                provideHttpClientTesting()
            ]

        }); 
        logService = TestBed.inject(LogService); 
        httpTesting = TestBed.inject(HttpTestingController);
    });

    afterEach(()=>{

        httpTesting.verify(); 

    }); 

    it('returns empty list of logs',()=>{

        const logs:Log[] = []; 
        logService.getLogs().subscribe(returnStatus => {

            expect(returnStatus.length).toBe(0);
            expect(returnStatus).toEqual([]); 
            
        });
        const request = httpTesting.expectOne('http://192.168.4.1/logs');
        expect(request.request.method).toBe("GET");
        request.flush(logs);

    }); 

    it('returns a list with one log',()=>{

        const log:Log = {
            name:"firstLog",
            date:'2026-07-26T11:18:29Z'
        }
        const logs:Log[] = [log]; 

        logService.getLogs().subscribe(returnStatus=>{

            expect(returnStatus.length).toBe(1);
            expect(returnStatus[0].name).toBe("firstLog");
        });
        const request = httpTesting.expectOne('http://192.168.4.1/logs');
        expect(request.request.method).toBe("GET");
        request.flush(logs);

    });

    it('returns a list with two logs',()=>{

        

    });
});