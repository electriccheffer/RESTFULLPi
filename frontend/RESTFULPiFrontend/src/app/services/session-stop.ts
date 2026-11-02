import { Injectable, Service } from '@angular/core';
import { Observable,of } from 'rxjs';
import { StopStatus } from '../generated/models/stop-status';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn:'root'
})
export class SessionStopService {

    protected url = "http://192.168.4.1:8080/logs/sessions"
    
    constructor(private http:HttpClient){}

    stopService(id:string):Observable<StopStatus>{
        let request:string = this.url + "/" + id + "/stop";
        return this.http.post<StopStatus>(request,{}); 
       
    }

}
