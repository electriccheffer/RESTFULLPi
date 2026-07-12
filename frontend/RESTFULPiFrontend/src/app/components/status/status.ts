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

  ngOnInit(){
    const data = this.statusService.getStatus();
    this.device = data.device; 
    this.status = data.status; 

  }
}
