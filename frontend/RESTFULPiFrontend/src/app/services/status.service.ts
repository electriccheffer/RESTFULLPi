import { Observable } from "rxjs"; 
import { Status } from "../generated/model/status";
import { inject, Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";

@Injectable({
    providedIn:'root',
})
export class StatusService{

    constructor(private http:HttpClient){}

    getStatus():Observable<Status>{

        return this.http.get("https://restfulpi.com/status");

    }

}