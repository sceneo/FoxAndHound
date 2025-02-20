import { Component } from '@angular/core';
import {CandidateApiService} from './candidate-api.service';
import {MatFormField} from '@angular/material/form-field';
import {MatOption, MatSelect} from '@angular/material/select';
import {CandidateDto} from './model/candidateDto';

@Component({
  selector: 'app-root',
  imports: [
    MatFormField,
    MatSelect,
    MatOption
  ],
  providers: [CandidateApiService],
  templateUrl: './hr.component.html',
  standalone: true,
  styleUrl: './hr.component.scss'
})
export class HrComponent {

  candidates: CandidateDto[] = [
    {userId: "UserIdHardcoded_1"},
    {userId: "UserIdHardcoded_2"}
  ]

  constructor(private candidateApiService: CandidateApiService) {

    this.candidateApiService.getCandidates()
      .subscribe((candidates) => {
      console.log(candidates);
    });

  }


}
