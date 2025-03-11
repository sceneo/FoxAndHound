import {Component, OnInit} from '@angular/core';
import {MatFormField, MatLabel} from '@angular/material/form-field';
import {MatOption} from '@angular/material/select';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatAutocomplete, MatAutocompleteSelectedEvent, MatAutocompleteTrigger} from '@angular/material/autocomplete';
import {MatInput} from '@angular/material/input';
import {map, Observable, of, startWith, switchMap, tap} from 'rxjs';
import {
  HeadDataService, ManagementAverageService,
  ManagementSummaryService,
  ModelsHeadDataDTO, ModelsManagementAverageDTO,
  ModelsManagementSummaryDTO
} from '../api';
import {CommonModule} from '@angular/common';

@Component({
  selector: 'app-root',
  imports: [
    CommonModule,
    MatFormField,
    MatOption,
    FormsModule,
    MatAutocomplete,
    MatAutocompleteTrigger,
    MatInput,
    MatLabel,
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
  managementAverage: ModelsManagementAverageDTO[] = [];
  filteredCandidates: Observable<string[]> = of([]);
  currentCandidate: ModelsManagementSummaryDTO | null = null;
  allHeadData: ModelsHeadDataDTO[] = [];
  currentHeadData: ModelsHeadDataDTO | undefined = undefined;
  private selectedUserMail: string | null = null;

  constructor(private headDataService: HeadDataService, private managementService: ManagementSummaryService, private managementAverageService: ManagementAverageService) {
  }
  ngOnInit() {
    this.isLoading = true;

    this.managementAverageService.managementAverageGet().subscribe( average => {
        this.managementAverage = average;
      })

    this.filteredCandidates = this.headDataService.managementAgreedCandidatesGet().pipe(
      tap(headDataSets => {
        this.allHeadData = headDataSets;
      }),
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
    )
  }

  private _filter(value: string): string[] {
    if (!value) {
      return [...this.candidates];
    }
    const filterValue = value.toLowerCase();
    return this.candidates.filter(option => option.toLowerCase().includes(filterValue));
  }

  onCandidateChange(event: MatAutocompleteSelectedEvent): void {
    this.selectedUserMail = String(event.option.value);
    this.currentHeadData = this.allHeadData.find(headData => headData.userEmail === this.selectedUserMail);
    this.loadCandidateData(this.selectedUserMail);
  }

  private loadCandidateData(email: string) {
    this.managementService.managementSummaryGet(email)
      .subscribe(managementSummary => {
        this.currentCandidate = managementSummary;
    })
  }

  showData(): boolean {
    return this.selectedUserMail !== "" && this.currentHeadData !== undefined;
  }

  namePrint(): string {
    return this.currentHeadData?.name || "";
  }

  agePrint(): string {
    return String(this.currentHeadData?.age) || "";
  }

  experiencePrint(): string {
    return this.currentHeadData?.experienceSince || "";
  }

  prodynaStartPrint(): string {
    return this.currentHeadData?.startAtProdyna || "";
  }

  abstractPrint(): string {
    return this.currentHeadData?.abstract || "";
  }
}
