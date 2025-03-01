import {Component, OnInit} from '@angular/core';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {CommonModule} from '@angular/common';
import {RatingSliderComponent} from '../rating-slider/rating-slider.component';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatIconModule} from '@angular/material/icon';
import { ModelsEmployerRatingDTO, RatingEmployerService } from '../api';
import { MatSelectChange, MatSelectModule } from '@angular/material/select';

@Component({
  selector: 'app-hr',
  templateUrl: './hr.component.html',
  styleUrls: ['./hr.component.scss'],
  imports: [
    CommonModule,
    MatIconModule,
    MatButtonModule,
    MatInputModule,
    MatFormFieldModule,
    MatSelectModule,
    ReactiveFormsModule,
    FormsModule,
    RatingSliderComponent
  ],
  standalone: true
})
export class HrComponent implements OnInit {
  isLoading: boolean = false;
  employerRatings: ModelsEmployerRatingDTO[] = [];
  candidates: string[] = [];
  categories: string[] = [];

  ratingForm: FormGroup | null = null;
  candidatesForm: FormGroup  = new FormGroup({
    "userEmail": new FormControl()
  });

  private selectedUserMail: string | null = null;

  constructor(private ratingService: RatingEmployerService) {}

  ngOnInit() {
    this.isLoading = true;
    
    this.ratingService.ratingsEmployerCandidatesGet().subscribe(users => {
      this.candidates = users;
      this.isLoading = false;
    });
  }

  onCandidateChange(event: MatSelectChange): void {
    this.selectedUserMail = String(event.value);

    this.ratingForm = null;
    this.isLoading = true;

    this.ratingService.ratingsEmployerGet(this.selectedUserMail)
      .subscribe(ratings => {
        this.categories = [...new Set(ratings.map(rating => String(rating.category)))];

        const controls: Record<string, FormControl> = {};

        ratings.forEach((rating) => {
          controls[`response_candidate_${rating.ratingCardId}`] = new FormControl({ value: rating.textResponseCandidate || "", disabled: true }, []);
          controls[`rating_candidate_${rating.ratingCardId}`] = new FormControl({ value: rating.ratingCandidate || 0, disabled: true }, []);
          controls[`response_employer_${rating.ratingCardId}`] = new FormControl(rating.textResponseEmployer || "", [Validators.required]);
          controls[`rating_employer_${rating.ratingCardId}`] = new FormControl(rating.ratingEmployer || 0, [Validators.required]);
        });
        this.ratingForm = new FormGroup(controls);
        this.isLoading = false;
        this.employerRatings = ratings;
    });
  }

  getRatingsOfCategory(category: string): ModelsEmployerRatingDTO[] {
    return this.employerRatings.filter(
      rating => rating.category === category
    );
  }

  onSubmit(): void {
    const formValues = this.ratingForm?.getRawValue();

    const updatedRatings: ModelsEmployerRatingDTO[] = this.employerRatings.map(rating => ({
      ...rating,
      textResponseEmployer: formValues[`response_employer_${rating.ratingCardId}`],
      ratingEmployer: formValues[`rating_employer_${rating.ratingCardId}`]
    }));

    this.isLoading = true;

    this.ratingService.ratingsEmployerPost(updatedRatings)
      .subscribe(() => this.isLoading = false);
  }
}
