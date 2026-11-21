import { HttpTestingController, provideHttpClientTesting } from "@angular/common/http/testing";
import { SessionRestartService } from "./session-restart.service";
import { TestBed } from "@angular/core/testing";
import { provideHttpClient } from "@angular/common/http";
import { RestartStatus } from "../generated/models/stop-status";

describe("SessionRestartService",() => {

    let service:SessionRestartService; 
    let httpMock:HttpTestingController; 

    beforeEach(() => {
        
        TestBed.configureTestingModule({

            providers: [
                SessionRestartService,
                provideHttpClient(),
                provideHttpClientTesting()
            ]

        }); 
        service = TestBed.inject(SessionRestartService); 
        httpMock = TestBed.inject(HttpTestingController); 
    }); 

    afterEach(() => {
        httpMock.verify();
    });

    it("Should create", () => {

        expect(service).toBeTruthy(); 

    }); 

    it("Should send POST /logs/sessions/{id}/restart", () => {
        
        let response:RestartStatus = {success:true}; 
        const expected:RestartStatus = {success:true}; 
        response.success = true; 
        const id:string = "xyz"; 
        service.restartSession(id).subscribe((res) => {
            expect(res).toEqual(expected); 
            expect(res.success).toBe(true); 
        }); 
        
        const request = httpMock.expectOne('http://192.168.4.1:8080/logs/sessions/xyz/restart'); 
        expect(request.request.method).toBe('POST'); 
        expect(request.request.body).toEqual({}); 
        request.flush(response); 
    }); 

}); 