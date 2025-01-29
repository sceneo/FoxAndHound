import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './admin.component.html',
  standalone: true,
  styleUrl: './admin.component.scss'
})
export class AdminComponent {
}
