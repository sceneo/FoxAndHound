import {Component, OnInit} from '@angular/core';
import {MatFormField, MatLabel} from '@angular/material/form-field';
import {MatOption, MatSelect} from '@angular/material/select';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatAutocomplete, MatAutocompleteSelectedEvent, MatAutocompleteTrigger} from '@angular/material/autocomplete';
import {MatButton} from '@angular/material/button';
import {MatInput} from '@angular/material/input';
import {MatSlideToggle} from '@angular/material/slide-toggle';
import {map, Observable, of, startWith, switchMap, tap} from 'rxjs';
import {HeadDataService, ManagmentSummaryService} from '../api';
import {CommonModule} from '@angular/common';

@Component({
  selector: 'app-root',
  imports: [
    CommonModule,
    MatFormField,
    MatSelect,
    MatOption,
    FormsModule,
    MatAutocomplete,
    MatAutocompleteTrigger,
    MatButton,
    MatInput,
    MatLabel,
    MatSlideToggle,
    ReactiveFormsModule
  ],
  templateUrl: './management-dashboard.component.html',
  standalone: true,
  styleUrl: './management-dashboard.component.scss'
})
export class ManagementDashboardComponent implements OnInit {
  isLoading: boolean = false;

  candidatesForm: FormGroup  = new FormGroup({
    "userEmail": new FormControl("", [Validators.email, Validators.required]),
  });

  candidates: string[] = [];
  filteredCandidates: Observable<string[]> = of([]);

  private selectedUserMail: string | null = null;

  constructor(private headDataService: HeadDataService, private managementService: ManagmentSummaryService) {
  }
  ngOnInit() {
    this.isLoading = true;

    this.filteredCandidates = this.headDataService.managementAgreedCandidatesGet().pipe(
      tap(headDataSets => {
        this.candidates = headDataSets
          .map(headData => headData.userEmail || "");
        this.isLoading = false;
      }),
      switchMap(() => {
        const mailControl = this.candidatesForm.controls["userEmail"];
        return mailControl.valueChanges.pipe(
          startWith(''),
          map(value => this._filter(value || ''))
        );
      })
    );
  }

  private _filter(value: string): string[] {
    if (!value) {
      return [...this.candidates];
    }
    const filterValue = value.toLowerCase();
    return this.candidates.filter(option => option.toLowerCase().includes(filterValue));
  }

  onCandidateChange(event: MatAutocompleteSelectedEvent): void {
    this.loadCandidateData(String(event.option.value));
  }

  private loadCandidateData(email: String) {
    window.alert("load " + email); // TODO
  }
}
