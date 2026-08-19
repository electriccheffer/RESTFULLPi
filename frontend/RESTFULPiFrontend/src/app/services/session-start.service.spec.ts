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

    })
});

