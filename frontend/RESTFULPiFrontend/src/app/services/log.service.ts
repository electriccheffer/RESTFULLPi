import { Observable } from "rxjs";
import { Log } from "../generated/model/log";

export abstract class LogService{

    abstract getLogs():Observable<Log[]>; 

}