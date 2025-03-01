import {ChangeDetectionStrategy, Component} from '@angular/core';
import {CandidateApiService} from './candidate-api.service';
import {MatFormField, MatFormFieldModule} from '@angular/material/form-field';
import {MatOption, MatSelect} from '@angular/material/select';
import {MatButton, MatButtonModule} from '@angular/material/button';
import {MatInput, MatInputModule} from '@angular/material/input';
import {CommonModule, NgForOf} from '@angular/common';
import {RatingSliderComponent} from '../rating-slider/rating-slider.component';
import {FormGroup, FormsModule, ReactiveFormsModule} from '@angular/forms';
import {MatIconModule} from '@angular/material/icon';

@Component({
  selector: 'app-root',
  imports: [
    CommonModule,
    MatIconModule,
    MatButtonModule,
    MatInputModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    FormsModule,
    RatingSliderComponent,
    MatSelect,
    MatOption
  ],
  providers: [CandidateApiService],
  templateUrl: './hr.component.html',
  standalone: true,
  styleUrl: './hr.component.scss'
})
export class HrComponent {

  /*candidates: CandidateDto[] = [
    {email: "test@test.com"},
    {email: "test@anotherTest.ch"}
  ]

  isLoading: boolean = false;
  isFormValid: boolean = false;
  ratingForm: FormGroup = new FormGroup({});

  private ratingCards: RatingCard[] = [];

  constructor(private candidateApiService: CandidateApiService, private ratingCardApiService: RatingCardApiService) {

    this.candidateApiService.getCandidates()
      .subscribe((candidates) => {
      console.log(candidates);
      // TODO: load candidates
    });

    this.ratingCardApiService
      .getRatingCards()
      .subscribe((ratingCardDtos: RatingCardDto[]) => {
        this.ratingCards = ratingCardDtos.map(dto => ({
          ...dto,
          rating: 0,
          response: ''
        }));
        this.isLoading = false;
      });

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

  updateRating(rating: number, cardId: string): void {
    const control = this.ratingForm.get(`rating_${cardId}`);
    if (control) {
      control.setValue(rating);
    }
  }

  onSubmit(): void {
    if (this.isFormValid) {
      window.alert('Form submitted from HR side');
      console.log(this.ratingForm?.value);  // Logs all form values (responses and ratings)
    } else {
      window.alert('Form is invalid.');
    }
  }*/


}
