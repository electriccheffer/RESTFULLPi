import { Observable, of, Subject, tap, throwError } from "rxjs";
import { StartStatus } from "../generated/models";
import { HttpClient, HttpErrorResponse, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { catchError } from "rxjs";

@Injectable({

    providedIn:'root'

})

export class SessionStartService{

    protected readonly requestUrl = 'http://192.168.4.1:8080/logs/sessions'; 
    protected readonly sessionStartedSubject = new Subject<void>(); 
    readonly sessionStarted$ = this.sessionStartedSubject.asObservable(); 

    constructor(private http:HttpClient){}

    startSession():Observable<StartStatus>{

        return this.http.post<StartStatus>(this.requestUrl,{}).pipe(tap(()=> {
            this.emitSessionStarted();}),
            catchError(this.handleError)); 

    };

    protected emitSessionStarted():void{

        this.sessionStartedSubject.next();
    }

    private handleError(error:HttpErrorResponse):Observable<never>{
        return throwError(() => error); 

    }

}