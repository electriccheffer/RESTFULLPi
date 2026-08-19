import { Observable, of, throwError } from "rxjs";
import { StartStatus } from "../generated/models";
import { HttpClient, HttpErrorResponse, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { catchError } from "rxjs";

@Injectable({

    providedIn:'root'

})

export class SessionStartService{

    private readonly requestUrl = 'http://192.168.4.1:8080/logs/sessions'; 

    constructor(private http:HttpClient){}

    startSession():Observable<StartStatus>{

        return this.http.post<StartStatus>(this.requestUrl,{}).pipe(catchError(this.handleError)); 

    };

    private handleError(error:HttpErrorResponse):Observable<never>{
        
        return throwError(() => error); 

    }

}