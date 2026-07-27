import { Observable } from "rxjs";
import { Log } from "../generated/model/log";
import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";

@Injectable({

    providedIn:'root', 

})
export class LogService{

    constructor(private http:HttpClient){}

    getLogs():Observable<Log[]>{

        return this.http.get<Log[]>("http://192.168.4.1/logs");

    } 

}