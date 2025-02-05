import { Component } from '@angular/core';
import {FormArray, FormBuilder, FormGroup, ReactiveFormsModule} from '@angular/forms';
import { CategoryArrangement } from '../models/category-arrangement.interface';
import { ContentService } from '../content/content.service';
import { RatingCardApiService } from './rating-card-api.service';
import { RatingCard } from '../models/rating-card.interface';
import { Category } from '../models/category.enum';
import {NgForOf, NgIf} from '@angular/common';
import {MatIcon} from '@angular/material/icon';
import {MatFormField} from '@angular/material/form-field';
import {MatButton} from '@angular/material/button';
import {MatInput} from '@angular/material/input';

@Component({
  selector: 'app-root',
  templateUrl: './senior-candidate.component.html',
  styleUrls: ['./senior-candidate.component.scss'],
  imports: [
    NgForOf,
    ReactiveFormsModule,
    NgIf,
    MatIcon,
    MatFormField,
    MatButton,
    MatInput,
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

    this.ratingCardApiService.getRatingCards()
      .subscribe(ratingCards => {
      this.ratingCards = ratingCards;
      this.fillRatingCardsToForm();
    },
        () =>  console.log("we can route to error here?") // TODO
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
      textResponse: [''],
      rating: [0]
    });
  }

  getArrangements(): CategoryArrangement[] {
    return ContentService.getCategoryArrangement();
  }

  get ratingCardForms(): FormArray {
    return this.form.get('ratingCards') as FormArray;
  }

  clearForm() {
    this.form.reset();
    this.ratingCardForms.clear();
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
    return this.ratingCardForms.controls.filter(card => card.value.category === category) as FormGroup[];
  }
}
