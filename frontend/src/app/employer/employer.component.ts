import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './employer.component.html',
  standalone: true,
  styleUrl: './employer.component.scss'
})
export class EmployerComponent {
  title = 'frontend';
}
