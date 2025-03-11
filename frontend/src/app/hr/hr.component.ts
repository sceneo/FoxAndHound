import {Component, OnInit} from '@angular/core';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {CommonModule} from '@angular/common';
import {RatingSliderComponent} from '../rating-slider/rating-slider.component';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatIconModule} from '@angular/material/icon';
import {HeadDataService, ModelsEmployerRatingDTO, ModelsHeadDataDTO, RatingEmployerService} from '../api';
import { MatSelectModule } from '@angular/material/select';
import { MatAutocompleteModule, MatAutocompleteSelectedEvent } from '@angular/material/autocomplete';
import { map, Observable, of, startWith, switchMap, tap } from 'rxjs';
import {MatSlideToggle} from '@angular/material/slide-toggle';
import {RatingGraphicComponent} from '../rating-graphic/rating-graphic.component';
import {MatCheckbox} from '@angular/material/checkbox';

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
    MatAutocompleteModule,
    MatSelectModule,
    ReactiveFormsModule,
    FormsModule,
    RatingSliderComponent,
    MatSlideToggle,
    RatingGraphicComponent,
    MatCheckbox
  ],
  standalone: true
})
export class HrComponent implements OnInit {
  isLoading: boolean = false;
  employerRatings: ModelsEmployerRatingDTO[] = [];
  candidates: string[] = [];
  categories: string[] = [];

  headDataForm: FormGroup | null = null;
  ratingForm: FormGroup | null = null;
  candidatesForm: FormGroup  = new FormGroup({
    "userEmail": new FormControl("", [Validators.email, Validators.required]),
    "hideCandidateInfo": new FormControl(true, []),
    "hideRating": new FormControl(true, [])
  });
  filteredCandidates: Observable<string[]> = of([]);

  private selectedUserMail: string | null = null;

  constructor(private ratingService: RatingEmployerService, private headDataService: HeadDataService) {}

  ngOnInit() {
    this.isLoading = true;

    this.filteredCandidates = this.ratingService.ratingsEmployerCandidatesGet().pipe(
      tap(users => {
        this.candidates = users;
        this.isLoading = false;
      }),
      switchMap(() => {
        const mailControl = this.candidatesForm.controls["userEmail"];
        return mailControl.valueChanges.pipe(
          tap(value => this.onMailTextChange(value)),
          startWith(''),
          map(value => this._filter(value || ''))
        );
      })
    );
  }

  onCandidateChange(event: MatAutocompleteSelectedEvent): void {
    this._initializeCandidate(String(event.option.value));
  }

  getRatingsOfCategory(category: string): ModelsEmployerRatingDTO[] {
    return this.employerRatings.filter(
      rating => rating.category === category
    );
  }

  save(): void {
    const formValues = this.ratingForm?.getRawValue();

    const updatedRatings: ModelsEmployerRatingDTO[] = this.employerRatings.map(rating => ({
      ...rating,
      textResponseEmployer: formValues[`response_employer_${rating.ratingCardId}`],
      ratingEmployer: formValues[`rating_employer_${rating.ratingCardId}`],
      notApplicableEmployer: formValues[`not_applicable_employer_${rating.ratingCardId}`],
    }));

    this.isLoading = true;
    const headDataFormValues = this.headDataForm?.getRawValue();

    const updateHeadData: ModelsHeadDataDTO = {
      abstract: headDataFormValues['abstract'],
      age: headDataFormValues['age'],
      agreedOn: headDataFormValues['agreedOn'],
      experienceSince: undefined, //headDataFormValues['experienceSince'],TODO decide on date format
      name: headDataFormValues['name'],
      startAtProdyna: undefined, // headDataFormValues['startAtProdyna'], TODO decide on date format
      userEmail: this.selectedUserMail ?? "",
    }

    this.headDataService.headDataPost(updateHeadData).subscribe(() => this.isLoading = false);
    this.ratingService.ratingsEmployerPost(updatedRatings).subscribe(() => this.isLoading = false);
  }

  private onMailTextChange(mailAddress: string): void {
    const mailControl = this.candidatesForm.controls["userEmail"];

    if (mailControl.valid && mailAddress !== this.selectedUserMail) {
      this._initializeCandidate(mailAddress);
    } else if (mailControl.invalid) {
      this.ratingForm = null;
    }
  }

  private _initializeCandidate(mailAddress: string): void {
    this.selectedUserMail = mailAddress;

    this.ratingForm = null;
    this.isLoading = true;

    this.headDataService.headDataGet(this.selectedUserMail).subscribe(headData => {
      this.headDataForm = new FormGroup({
        "agreedOn": new FormControl(headData.agreedOn || false, undefined),
        "name": new FormControl(headData.name || "", []),
        "experienceSince": new FormControl(headData.experienceSince || "", []),
        "startAtProdyna": new FormControl(headData.startAtProdyna || "", []),
        "age": new FormControl(headData.age || 0, []),
        "abstract": new FormControl(headData.abstract ||"", undefined),
      });
    },
      () => {
        this.headDataForm = new FormGroup({
          "agreedOn": new FormControl(false, undefined),
          "name": new FormControl("", []),
          "experienceSince": new FormControl( "", []),
          "startAtProdyna": new FormControl( "", []),
          "age": new FormControl(0, []),
          "abstract": new FormControl("", undefined),
        });
      })




    this.ratingService.ratingsEmployerGet(this.selectedUserMail)
      .subscribe(ratings => {
        this.categories = [...new Set(ratings.map(rating => String(rating.category)))];

        const controls: Record<string, FormControl> = {};

        // TODO: we need to load the average rating. currently we randomly pick a number
        ratings.forEach((rating) => {
          controls[`response_candidate_${rating.ratingCardId}`] = new FormControl({ value: rating.textResponseCandidate || "", disabled: true }, []);
          controls[`rating_candidate_${rating.ratingCardId}`] = new FormControl({ value: rating.ratingCandidate || 0, disabled: true }, []);
          controls[`not_applicable_candidate_${rating.ratingCardId}`] = new FormControl({ value: rating.notApplicableCandidate || false, disabled: true }, []);
          controls[`response_employer_${rating.ratingCardId}`] = new FormControl(rating.textResponseEmployer || "", [Validators.required]);
          controls[`rating_employer_${rating.ratingCardId}`] = new FormControl(rating.ratingEmployer || 0, [Validators.required]);
          controls[`not_applicable_employer_${rating.ratingCardId}`] = new FormControl(rating.notApplicableEmployer || false, undefined);
          controls[`average_rating_${rating.ratingCardId}`] = new FormControl(Math.random() * 120, []);
        });
        this.ratingForm = new FormGroup(controls);
        this.isLoading = false;
        this.employerRatings = ratings;
    });
  }

  private _filter(value: string): string[] {
    if (!value) {
      return [...this.candidates];
    }
    const filterValue = value.toLowerCase();
    return this.candidates.filter(option => option.toLowerCase().includes(filterValue));
  }

  hideCandidateAnswers(): boolean {
    return this.candidatesForm.get('hideCandidateInfo')?.value ?? true;
  }

  hideRatings(): boolean {
    return this.candidatesForm.get('hideRating')?.value ?? true;
  }

  getCurrentRating(ratingCardId: number | undefined): number {
    if(ratingCardId) {
      return this.ratingForm?.get('rating_employer_' + ratingCardId)?.value ?? 0;
    }
    return 0;
  }

  getComparisonRating(ratingCardId: number | undefined): number {
    if(ratingCardId) {
      return this.ratingForm?.get('average_rating_' + ratingCardId)?.value ?? 100.0;

    }
    return 0;
  }

  isSaveButtonEnabled(): boolean {
    return this.headDataForm !== null && this.candidatesForm !== null;
  }
}
