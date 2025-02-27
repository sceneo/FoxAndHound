import {Component, OnInit} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {RatingSliderComponent} from '../rating-slider/rating-slider.component'; // Reactive forms
import { ModelsCandidateRatingDTO, RatingService } from '../api';

@Component({
  selector: 'app-root',
  templateUrl: './senior-candidate.component.html',
  styleUrls: ['./senior-candidate.component.scss'],
  imports: [
    CommonModule,
    MatIconModule,
    MatButtonModule,
    MatInputModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    FormsModule,
    RatingSliderComponent
  ],
  standalone: true
})
export class SeniorCandidateComponent implements OnInit {
  isLoading: boolean = false;
  candidateRatings: ModelsCandidateRatingDTO[] = [];
  ratingForm!: FormGroup;
  categories: string[] = [];

  constructor(private ratingService: RatingService) {}

  ngOnInit() {
    this.isLoading = true;
    // TODO: Here the mail of the logged in user needs to be set... since no auth till yet auto set to next senior :-)
    this.ratingService.ratingsCandidateGet("thomas.lederer@prodyna.com").subscribe((ratingCardDtos: ModelsCandidateRatingDTO[]) => {
      this.candidateRatings = ratingCardDtos;
      this.categories = [...new Set(ratingCardDtos.map(rating => String(rating.category)))];
      this.isLoading = false;
      this.initializeForm();
    });
  }

  initializeForm() {
    const controls: Record<string, FormControl> = {};
    controls[`email`] = new FormControl("", [Validators.required, Validators.email]);
    this.candidateRatings.forEach((rating) => {
      controls[`response_${rating.ratingCardId}`] = new FormControl(rating.textResponseCandidate || "", [Validators.required]);
      controls[`rating_${rating.ratingCardId}`] = new FormControl(rating.ratingCandidate || 0, [Validators.required]);
    });

    this.ratingForm = new FormGroup(controls);  // Assign FormGroup after initialization

    // TODO: use unsubscribe
  }


  getRatingsOfCategory(category: string): ModelsCandidateRatingDTO[] {
    return this.candidateRatings.filter(
      rating => rating.category === category
    );
  }

  onSubmit(): void {
    const formValues = this.ratingForm.getRawValue();
  
    const updatedRatings: ModelsCandidateRatingDTO[] = this.candidateRatings.map(rating => ({
      ...rating,
      textResponseCandidate: formValues[`response_${rating.ratingCardId}`],
      ratingCandidate: formValues[`rating_${rating.ratingCardId}`]
    }));

    this.ratingService.ratingsCandidatePost(updatedRatings).subscribe();
  }
}
