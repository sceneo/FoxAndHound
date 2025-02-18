import { Component } from '@angular/core';
import {CandidateApiService} from './candidate-api.service';

@Component({
  selector: 'app-root',
  imports: [],
  providers: [CandidateApiService],
  templateUrl: './employer.component.html',
  standalone: true,
  styleUrl: './employer.component.scss'
})
export class EmployerComponent {

  constructor(private candidateApiService: CandidateApiService) {

    this.candidateApiService.getCandidates()
      .subscribe((candidates) => {
      console.log(candidates);
    });

  }


}
