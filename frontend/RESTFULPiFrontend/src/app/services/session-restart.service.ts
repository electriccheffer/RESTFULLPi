import { catchError, Observable,throwError } from "rxjs";
import { RestartStatus } from "../generated/models/stop-status";
import { Injectable, InjectionToken } from "@angular/core";
import { HttpClient, HttpErrorResponse } from "@angular/common/http";

export abstract class RestartService{

    abstract restartSession(id:string):Observable<RestartStatus>;

}

@Injectable({
    providedIn: 'root'
})
export class SessionRestartService extends RestartService{
    
    private url:string = "http://192.168.4.1:8080/logs/sessions/"
    constructor(private http:HttpClient){super();}

    override restartSession(id: string): Observable<RestartStatus> {
        let request:string = this.url + id + "/restart"; 
        return this.http.post<RestartStatus>(request,{}).pipe(
                catchError((error:HttpErrorResponse) => {
                    return throwError(() => error); 
                })
            );
    }
}
