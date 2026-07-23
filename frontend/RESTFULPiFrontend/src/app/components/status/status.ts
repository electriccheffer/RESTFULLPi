import { Component } from '@angular/core';
import { StatusService } from '../../services/status.service';

@Component({
  selector: 'status',
  imports: [],
  templateUrl: './status.html',
  styleUrl: './status.css',
})
export class StatusComponent {
  
  device = ''; 
  status = ''; 

  constructor(private statusService:StatusService){}

  ngOnInit():void{
    this.statusService.getStatus().subscribe({
      next: (data) => {
        this.device = data.device ?? ''; 
        this.status = data.status ?? ''; 
      },
      error:() => {

        this.device = 'RESTFULPi';
        this.status = 'Server Unavailable, Unknown Status';         

      },
    });
    

  }
}
