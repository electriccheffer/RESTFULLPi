import { Component, Input } from '@angular/core';
import { SessionStopService } from '../../services/session-stop';
import { Log } from '../../generated/models';
import { StopStatus } from '../../generated/models/stop-status';

@Component({
  selector: 'app-session-stop',
  imports: [],
  templateUrl: './session-stop.html',
  styleUrl: './session-stop.css',
})
export class SessionStop {
  
  status?:StopStatus;
  errorMessage?:string; 

  @Input() log!:Log; 

  constructor(private service:SessionStopService){}

  stopSession():void{
      this.status = this.service.stopService(this.log.id!).
                          subscribe({

                      next:(data)=>{
                        this.status = data; 
                      },
                      error:(err)=>{
                        this.errorMessage = err.statusText;                           
                      }
                                  });
  } 
}
