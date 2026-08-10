import { Observable } from "rxjs"; 
import { Status } from "../generated/model/status";
import {Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";

@Injectable({
    providedIn:'root',
})
export class StatusService{

    constructor(private http:HttpClient){}

    getStatus():Observable<Status>{

        return this.http.get<Status>("http://192.168.4.1:8080/status");

    }

}