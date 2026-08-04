import { Component } from '@angular/core';
import { StatusComponent } from '../status/status';
import { Logs } from '../logs/logs';

@Component({
  selector: 'app-home-page',
  imports: [StatusComponent,Logs],
  templateUrl: './home-page.html',
  styleUrl: './home-page.css',
})
export class HomePage {}
