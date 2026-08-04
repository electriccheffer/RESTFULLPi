import { Component } from '@angular/core';
import { StatusService } from '../../services/status.service';

@Component({
  selector: 'status',
  imports: [],
  templateUrl: './status.html',
  styleUrl: './status.css',
})
export class StatusComponent {
  
  device = 'RESTFULPi'; 
  status = 'Server Unavailable, Unknown Status'; 

  constructor(private statusService:StatusService){}

  ngOnInit():void{
    this.statusService.getStatus().subscribe({
      next: (data) => {
        this.device = data?.device || 'RESTFULPi'; 
        this.status = data?.status || 'Server Unavailable, Unknown Status'; 
      },
      error:() => {

        this.device = 'RESTFULPi';
        this.status = 'Server Unavailable, Unknown Status';         

      },
    });
    

  }
}
