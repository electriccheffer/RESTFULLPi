import { Component, signal } from '@angular/core';
import { StatusService } from '../../services/status.service';

@Component({
  selector: 'status',
  imports: [],
  templateUrl: './status.html',
  styleUrl: './status.css',
})
export class StatusComponent {
  
  device = signal<string>('RESTFULPi'); 
  status = signal<string>('Server Unavailable, Unknown Status'); 

  constructor(private statusService:StatusService){}

  ngOnInit():void{
    this.statusService.getStatus().subscribe({
      next: (data) => {
       
        this.device.set(data?.device || 'RESTFULPi'); 
        this.status.set(data?.status || 'Server Unavailable, Unknown Status'); 
      },
      error:() => {
      
        this.device.set('RESTFULPi');
        this.status.set('Server Unavailable, Unknown Status');         

      },
    });
    

  }
}
