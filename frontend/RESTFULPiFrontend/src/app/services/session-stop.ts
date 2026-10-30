import { Injectable, Service } from '@angular/core';
import { Observable,of } from 'rxjs';
import { StopStatus } from '../generated/models/stop-status';

@Injectable({
    providedIn:'root'
})
export class SessionStopService {

    stopService():Observable<StopStatus>{

        return of({success:true}); 
    }

}
