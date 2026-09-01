import { ChangeDetectionStrategy, Component, OnDestroy, signal } from '@angular/core';
import { Log } from '../../generated/models/log';
import { LogService } from '../../services/log.service';
import { CommonModule } from '@angular/common';
import { single, Subject, takeUntil } from 'rxjs';
import { SessionStartService } from '../../services/session-start.service';

@Component({
  selector: 'app-logs',
  imports: [CommonModule],
  templateUrl: './logs.html',
  styleUrl: './logs.css',
  
})
export class Logs implements OnDestroy {

  logs = signal<Log[]>([]); 
  
  private readonly destroy$ = new Subject<void>(); 

  constructor(private logService:LogService,private sessionStartService:SessionStartService){}

  ngOnInit():void{

    this.getLogs();
    this.sessionStartService.sessionStarted$.pipe(takeUntil(this.destroy$)).    
      subscribe(() => {
        this.getLogs();
      });
  }

  getLogs():void{

    this.logService.getLogs().subscribe({
      next:(data) => {
        this.logs.set(data ?? []);
      }, error: () => {
        this.logs.set([]); 
      }
    }); 

  }

  ngOnDestroy():void{

    this.destroy$.next();
    this.destroy$.complete();

  }

}
