import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './unauthorized.component.html',
  standalone: true,
  styleUrl: './unauthorized.component.scss'
})
export class UnauthorizedComponent {
}
