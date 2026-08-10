import { TestBed } from "@angular/core/testing";
import { provideHttpClient } from "@angular/common/http";
import { HttpTestingController,provideHttpClientTesting } from "@angular/common/http/testing";
import { StatusService } from "./status.service";
import {Status} from "../generated/model/status";

describe('StatusService' ,() => {

    let service:StatusService; 
    let httpTesting:HttpTestingController;

    beforeEach(() => {

        TestBed.configureTestingModule({
            providers:[
                provideHttpClient(),
                provideHttpClientTesting()
            ]
        });
        
        service = TestBed.inject(StatusService);
        httpTesting = TestBed.inject(HttpTestingController);

    });

    afterEach(() => {

        httpTesting.verify();

    }); 

    it('test for online status of GET /status online case', () => {

        const backendStatus: Status = {

            device:'RESTFULPi',
            status:'Online'

        };

        service.getStatus().subscribe(returnStatus => {

            expect(returnStatus.device).toBe('RESTFULPi');
            expect(returnStatus.status).toBe('Online'); 
        }); 

        const request = httpTesting.expectOne('http://192.168.4.1:8080/status'); 
        expect(request.request.method).toBe('GET');
        request.flush(backendStatus);

    }); 

});