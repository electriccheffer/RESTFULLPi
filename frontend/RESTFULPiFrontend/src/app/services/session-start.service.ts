import { Observable, of } from "rxjs";
import { StartStatus } from "../generated/models";
import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";


@Injectable({

    providedIn:'root'

})

export class SessionStartService{

    private readonly requestUrl = 'http://192.168.4.1:8080/logs/sessions'; 

    constructor(private http:HttpClient){}

    startSession():Observable<StartStatus>{

        return this.http.post<StartStatus>(this.requestUrl,{}); 

    };

}