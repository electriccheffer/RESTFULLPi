import { ChangeDetectionStrategy, Component, signal } from '@angular/core';
import { Log } from '../../generated/model/log';
import { LogService } from '../../services/log.service';
import { CommonModule } from '@angular/common';
import { single } from 'rxjs';

@Component({
  selector: 'app-logs',
  imports: [CommonModule],
  templateUrl: './logs.html',
  styleUrl: './logs.css',
  
})
export class Logs {

  logs = signal<Log[]>([]); 

  constructor(private logService:LogService){}

  ngOnInit():void{

    this.logService.getLogs().subscribe({
      next:(data)=>{
        this.logs.set(data ?? []); 
      }, error: () => {
          
          this.logs = signal<Log[]>([]); 

      }

    });

  }

}
