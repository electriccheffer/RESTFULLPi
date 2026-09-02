import { ChangeDetectionStrategy, Component } from '@angular/core';
import { StatusComponent } from '../status/status';
import { Logs } from '../logs/logs';
import { SessionStart } from "../session-start/session-start";

@Component({
  selector: 'app-home-page',
  imports: [StatusComponent, Logs, SessionStart],
  templateUrl: './home-page.html',
  styleUrl: './home-page.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class HomePage {}
