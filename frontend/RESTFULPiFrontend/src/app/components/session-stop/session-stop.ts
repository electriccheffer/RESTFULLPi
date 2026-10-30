import { Component } from '@angular/core';
import { SessionStopService } from '../../services/session-stop';

@Component({
  selector: 'app-session-stop',
  imports: [],
  templateUrl: './session-stop.html',
  styleUrl: './session-stop.css',
})
export class SessionStop {

  constructor(private servic:SessionStopService){}

}
