import { Component } from '@angular/core';
import {FormArray, FormBuilder, FormGroup, ReactiveFormsModule} from '@angular/forms';
import {MatFormField} from '@angular/material/form-field';
import {MatButton} from '@angular/material/button';
import {MatInput} from '@angular/material/input';
import {CategoryArrangement} from '../models/category-arrangement.interface';
import {ContentService} from '../content/content.service';
import {Category} from '../models/category.enum';
import {NgForOf} from '@angular/common';
import {RatingCardApiService} from './rating-card-api.service';
import {RatingCard} from '../models/rating-card.interface';

@Component({
  selector: 'app-root',
  imports: [MatFormField, ReactiveFormsModule, MatButton, MatInput, NgForOf],
  providers: [RatingCardApiService],
  templateUrl: './senior-candidate.component.html',
  standalone: true,
  styleUrl: './senior-candidate.component.scss'
})
export class SeniorCandidateComponent {
  form: FormGroup;

  private ratingCards: RatingCard[] = [];

  constructor(private formBuilder: FormBuilder, ratingCardApiService: RatingCardApiService) {
    this.form = this.formBuilder.group({
      ratingCards: this.formBuilder.array([])
    });

    // TODO: make a loading animation
    ratingCardApiService.getRatingCards().subscribe(ratingCards => {
      this.ratingCards = ratingCards;
      this.fillRatingCardsToForm();
    });
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
      averageRating: [card.averageRating],
      orderId: [card.orderId],
      textResponse: [''],
      rating: [0]
    });
  }

  getArrangements(): CategoryArrangement[] {
    return ContentService.getCategoryArrangement();
  }

  getSortedCategoryCards(category: Category) {
    return this.ratingCardForms.controls.filter(ratingCard => ratingCard.value.category === category) // TODO: filter
  }

  get ratingCardForms(): FormArray {
    return this.form.get('ratingCards') as FormArray;
  }

  clearForm() {
    this.form.reset();
    this.ratingCardForms.clear();
  }

  saveForm() {
    if (this.form.valid) {
      console.log('Form saved:', this.form.value);
    } else {
      console.error('Form is invalid.');
    }
  }

  onSubmit() {
    if (this.form.valid) {
      console.log('Form submitted:', this.form.value);
    } else {
      console.error('Form is invalid.');
    }
  }
}
