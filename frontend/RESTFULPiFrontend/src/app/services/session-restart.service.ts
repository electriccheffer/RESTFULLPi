import { Observable } from "rxjs";
import { RestartStatus } from "../generated/models/stop-status";
import { Injectable, InjectionToken } from "@angular/core";


export abstract class RestartService{

    abstract restartSession(id:string):Observable<RestartStatus>;

}

