import { Component } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { MatFormField } from '@angular/material/form-field';
import { MatButton } from '@angular/material/button';
import { MatInput } from '@angular/material/input';
import { MatIcon } from '@angular/material/icon';
import { NgForOf, NgClass } from '@angular/common';
import { CategoryArrangement } from '../models/category-arrangement.interface';
import { ContentService } from '../content/content.service';
import { RatingCardApiService } from './rating-card-api.service';
import { RatingCard } from '../models/rating-card.interface';
import {Category} from '../models/category.enum';

@Component({
  selector: 'app-root',
  imports: [MatFormField, ReactiveFormsModule, MatButton, MatInput, MatIcon, NgForOf, NgClass],
  providers: [RatingCardApiService],
  templateUrl: './senior-candidate.component.html',
  standalone: true,
  styleUrls: ['./senior-candidate.component.scss']
})
export class SeniorCandidateComponent {
  form: FormGroup;
  private ratingCards: RatingCard[] = [];

  constructor(private formBuilder: FormBuilder, private ratingCardApiService: RatingCardApiService) {
    this.form = this.formBuilder.group({
      ratingCards: this.formBuilder.array([])
    });

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

  get ratingCardForms(): FormArray {
    return this.form.get('ratingCards') as FormArray;
  }

  clearForm() {
    this.form.reset();
    this.ratingCardForms.clear();
  }

  updateRating(index: number, rating: number) {
    const ratingControl = this.ratingCardForms.at(index).get('rating');
    if (ratingControl) {
      ratingControl.setValue(rating);
    }
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

  getPowerBarClass(rating: number, averageRating: number): string {
    const difference = rating - averageRating;
    if (difference <= -2) return 'red-bar';
    if (difference === -1) return 'orange-bar';
    if (difference === 0) return 'neutral-bar';
    if (difference === 1) return 'light-green-bar';
    return 'dark-green-bar';
  }
}
