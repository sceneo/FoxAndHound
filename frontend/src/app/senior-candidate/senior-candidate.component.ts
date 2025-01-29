import { Component } from '@angular/core';
import {FormArray, FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatFormField, MatLabel} from '@angular/material/form-field';
import {MatButton} from '@angular/material/button';
import {MatInput} from '@angular/material/input';

@Component({
  selector: 'app-root',
  imports: [MatFormField, ReactiveFormsModule, MatButton, MatInput, MatLabel],
  templateUrl: './senior-candidate.component.html',
  standalone: true,
  styleUrl: './senior-candidate.component.scss'
})
export class SeniorCandidateComponent {
  form: FormGroup;

  constructor(private fb: FormBuilder) {
    this.form = this.fb.group({
      ratingCards: this.fb.array([])
    });
  }

  getCategory1Cards() {
    return this.ratingCards.controls; // TODO: filter
  }

  get ratingCards(): FormArray {
    return this.form.get('ratingCards') as FormArray;
  }

  clearForm() {
    this.form.reset();
    this.ratingCards.clear();
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
