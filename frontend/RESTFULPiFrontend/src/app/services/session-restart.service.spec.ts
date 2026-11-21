import { HttpTestingController, provideHttpClientTesting } from "@angular/common/http/testing";
import { SessionRestartService } from "./session-restart.service";
import { TestBed } from "@angular/core/testing";
import { provideHttpClient } from "@angular/common/http";

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

    it("Should create", () => {

        expect(service).toBeTruthy(); 

    }); 

    
}); 