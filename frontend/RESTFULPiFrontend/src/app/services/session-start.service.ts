import { Observable } from "rxjs";
import { StartStatus } from "../generated/models";

export abstract class SessionStartService{

    abstract startSession():Observable<StartStatus>;

}