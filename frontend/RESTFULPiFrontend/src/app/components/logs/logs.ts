import { ChangeDetectionStrategy, Component } from '@angular/core';
import { Log } from '../../generated/model/log';
import { LogService } from '../../services/log.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-logs',
  imports: [CommonModule],
  templateUrl: './logs.html',
  styleUrl: './logs.css',
  
})
export class Logs {

  logs:Log[] = []; 

  constructor(private logService:LogService){}

  ngOnInit():void{

    this.logService.getLogs().subscribe({
      next:(data)=>{
        this.logs = data; 
      }, error: () => {
          
          this.logs = []; 

      }

    });

  }

}
