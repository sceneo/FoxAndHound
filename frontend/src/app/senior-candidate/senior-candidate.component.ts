import {Component, OnInit} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {ContentService} from '../content/content.service';
import {CategoryArrangement} from '../models/category-arrangement.interface';
import {Category} from '../models/category.enum';
import {RatingSliderComponent} from '../rating-slider/rating-slider.component'; // Reactive forms
import { ModelsRatingCard, RatingService } from '../api';

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
  ratingCards: ModelsRatingCard[] = [];
  ratingForm: FormGroup = new FormGroup({});
  isFormValid: boolean = false;

  constructor(private ratingCardService: RatingService) {}

  ngOnInit() {
    this.isLoading = true;
    this.ratingCardService.ratingCardsGet().subscribe((ratingCardDtos: ModelsRatingCard[]) => {
      this.ratingCards = ratingCardDtos;
      this.isLoading = false;
      this.initializeForm();
    });
  }

  initializeForm() {
    const controls: Record<string, FormControl> = {};
    controls[`email`] = new FormControl("", [Validators.required, Validators.email]);
    this.ratingCards.forEach((card) => {
      controls[`response_${card.id}`] = new FormControl("", [Validators.required]);
      controls[`rating_${card.id}`] = new FormControl(0, [Validators.required]);
    });

    this.ratingForm = new FormGroup(controls);  // Assign FormGroup after initialization

    // TODO: use unsubscribe
  }

  getArrangements(): CategoryArrangement[] {
    return ContentService.getCategoryArrangement().map(arrangement => {
      return {
        ...arrangement,
        ratingCards: this.getCategoryFilteredCards(arrangement.category)
      };
    }).filter(arrangement => arrangement.ratingCards.length > 0);
  }

  private getCategoryFilteredCards(category: Category): ModelsRatingCard[] {
    return this.ratingCards.filter(
      card => card.category === category
    );
  }

  onSubmit(): void {
    if (this.isFormValid) {
      window.alert('Form submitted from candidate side');
      console.log(this.ratingForm?.value);  // Logs all form values (responses and ratings)
    } else {
      window.alert('Form is invalid.');
    }
  }

  updateRating(rating: number, cardId: number | undefined): void {
    const control = this.ratingForm.get(`rating_${cardId}`);
    if (control) {
      control.setValue(rating);
    }
  }
}
