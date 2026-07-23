import { Observable } from "rxjs"; 
import { Status } from "../generated/model/status";
import { inject } from "@angular/core";
import { HttpClient } from "@angular/common/http";

export abstract class StatusService{

    abstract getStatus():Observable<Status>;

}