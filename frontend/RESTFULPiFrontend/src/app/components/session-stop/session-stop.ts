import { Component, Inject, Input } from '@angular/core';
import { SessionStopService } from '../../services/session-stop';
import { Log } from '../../generated/models';
import { RestartStatus, StopStatus } from '../../generated/models/stop-status';
import { RestartService } from '../../services/session-restart.service';

@Component({
  selector: 'app-session-stop',
  imports: [],
  templateUrl: './session-stop.html',
  styleUrl: './session-stop.css',
})
export class SessionStop {
  
  status?:StopStatus;
  errorMessage:string = ''; 
  restartStatus?:RestartStatus; 
 
  @Input() log!:Log; 

  constructor(private service:SessionStopService,
              private restartService:RestartService){}

  stopSession():void{
      this.service.stopService(this.log.id!).
                          subscribe({

                      next:(data)=>{
                        this.status = data; 
                      },
                      error:(err)=>{
                        this.errorMessage = err.statusText;                           
                      }
                                  });
  }

  startSession():void{

    this.restartService.restartSession(this.log.id!).subscribe({
        next:(data) => {
            this.restartStatus = data; 
        },
        error:(err)=>{
          this.errorMessage = err.statusText; 
        }
    });

  }
}
