import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './senior-candidate.component.html',
  standalone: true,
  styleUrl: './senior-candidate.component.scss'
})
export class SeniorCandidateComponent {
  title = 'frontend';
}
