import { Component } from '@angular/core';
import {FormArray, FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import { CategoryArrangement } from '../models/category-arrangement.interface';
import { ContentService } from '../content/content.service';
import { RatingCardApiService } from './rating-card-api.service';
import { RatingCard } from '../models/rating-card.interface';
import { Category } from '../models/category.enum';
import {CommonModule} from '@angular/common';
import { MatIconModule} from '@angular/material/icon';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';

@Component({
  selector: 'app-root',
  templateUrl: './senior-candidate.component.html',
  styleUrls: ['./senior-candidate.component.scss'],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatIconModule,
    MatFormFieldModule,
    MatButtonModule,
    MatInputModule,
  ],
  providers: [RatingCardApiService],
  standalone: true
})
export class SeniorCandidateComponent {
  form: FormGroup;
  private ratingCards: RatingCard[] = [];

  constructor(private formBuilder: FormBuilder, private ratingCardApiService: RatingCardApiService) {
    this.form = this.formBuilder.group({
      ratingCards: this.formBuilder.array([])
    });

    this.ratingCardApiService.getRatingCards().subscribe(
      (ratingCards) => {
        this.ratingCards = ratingCards;
        this.fillRatingCardsToForm();
      },
      () => console.log('Error loading rating cards') // Error handling if needed
    );
  }

  fillRatingCardsToForm(): void {
    const ratingCardsArray = this.form.get('ratingCards') as FormArray;
    this.ratingCards.forEach(card => {
      ratingCardsArray.push(this.createRatingCardFormGroup(card));
    });
  }

  createRatingCardFormGroup(card: RatingCard): FormGroup {
    return this.formBuilder.group({
      id: [card.id],
      question: [card.question],
      category: [card.category],
      orderId: [card.orderId],
      textResponse: ["", Validators.required],
      rating: [0, Validators.required]
    });
  }

  getArrangements(): CategoryArrangement[] {
    return ContentService.getCategoryArrangement();
  }

  get ratingCardForms(): FormArray {
    return this.form.get('ratingCards') as FormArray;
  }


  updateRating(cardControl: FormGroup, rating: number) {
    cardControl.get('rating')?.setValue(rating);
  }

  onSubmit() {
    if (this.form.valid) {

      console.log('Form submitted:', this.form.value);
    } else {
      console.error('Form is invalid.');
    }
  }

  getCategoryFilteredCards(category: Category): FormGroup[] {
    return this.ratingCardForms.controls.filter(
      card => card.value.category === category
    ) as FormGroup[];
  }
}
