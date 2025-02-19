import {ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {RatingCardApiService} from './rating-card-api.service';
import {RatingCardDto} from '../models/rating-card-dto.interface';
import {CommonModule} from '@angular/common';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSliderModule} from '@angular/material/slider';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {ContentService} from '../content/content.service';
import {CategoryArrangement} from '../models/category-arrangement.interface';
import {RatingCard} from '../models/rating-card.interface';
import {Category} from '../models/category.enum';
import {RatingSliderComponent} from '../rating-slider/rating-slider.component'; // Reactive forms

@Component({
  selector: 'app-root',
  templateUrl: './senior-candidate.component.html',
  styleUrls: ['./senior-candidate.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
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
  providers: [RatingCardApiService],
  standalone: true
})
export class SeniorCandidateComponent implements OnInit {
  isLoading: boolean = false;
  ratingCards: RatingCard[] = [];
  ratingForm: FormGroup = new FormGroup({});
  isFormValid: boolean = false;

  constructor(private ratingCardApiService: RatingCardApiService, private cdRef: ChangeDetectorRef) {}

  ngOnInit() {
  this.cdRef.detach();
  this.isLoading = true;
    this.ratingCardApiService
      .getRatingCards()
      .subscribe((ratingCardDtos: RatingCardDto[]) => {
        this.ratingCards = ratingCardDtos.map(dto => ({
          ...dto,
          rating: 0,
          response: ''
        }));
        this.isLoading = false;
          this.initializeForm();
          this.cdRef.detectChanges();
      });
  }

  initializeForm() {
    const controls: Record<string, FormControl> = {}; // Accumulator for controls

    this.ratingCards.forEach((card) => {
      controls[`response_${card.id}`] = new FormControl(card.response || "", [Validators.required]);
      controls[`rating_${card.id}`] = new FormControl(card.rating || 0, [Validators.required]);
    });

    this.ratingForm = new FormGroup(controls);  // Assign FormGroup after initialization

    // TODO: use unsubscribe
    /**
    this.ratingForm.valueChanges.subscribe(() => {
      this.isFormValid = this.ratingForm?.valid;
    })
      **/
  }

  getArrangements(): CategoryArrangement[] {
    return ContentService.getCategoryArrangement().map(arrangement => {
      return {
        ...arrangement,
        ratingCards: this.getCategoryFilteredCards(arrangement.category)
      };
    }).filter(arrangement => arrangement.ratingCards.length > 0);
  }

  private getCategoryFilteredCards(category: Category): RatingCard[] {
    return this.ratingCards.filter(
      card => card.category === category
    );
  }

  onSubmit(): void {
    if (this.isFormValid) {
      window.alert('Form submitted');
      console.log(this.ratingForm?.value);  // Logs all form values (responses and ratings)
    } else {
      window.alert('Form is invalid.');
    }
  }

  updateRating(rating: number, cardId: string): void {
    const control = this.ratingForm.get(`rating_${cardId}`);
    if (control) {
      control.setValue(rating);
    }
  }
}
