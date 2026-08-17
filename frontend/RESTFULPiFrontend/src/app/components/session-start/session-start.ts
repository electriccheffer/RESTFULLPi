import { Component } from '@angular/core';
import { Observable, of } from 'rxjs';
import { StartStatus } from '../../generated/models';
import { SessionStartService } from '../../services/session-start.service';

@Component({
  selector: 'app-session-start',
  imports: [],
  templateUrl: './session-start.html',
  styleUrl: './session-start.css',
})
export class SessionStart {

  startStatus?:StartStatus; 
  errorMessage?:string; 

  constructor(private startSessionService:SessionStartService){}

  ngOnInit():void{

    

  }

  startLogging():void{
    this.startStatus = undefined;
    this.errorMessage = undefined; 
    this.startSessionService.startSession().subscribe({
      next:(data)=>{
        this.startStatus = data; 
      },
      error: (err) => {
        
        this.errorMessage = 'Error starting file check the pi';

      }

    });
    

  }
}
