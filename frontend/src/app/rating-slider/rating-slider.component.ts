import {ChangeDetectorRef, Component, EventEmitter, Output} from '@angular/core';
import {MatSlider, MatSliderThumb} from '@angular/material/slider';
import {MatFormField, MatLabel} from '@angular/material/form-field';

@Component({
  selector: 'app-rating-slider',
  imports: [
    MatSlider,
    MatSliderThumb,
    MatLabel,
    MatFormField
  ],
  templateUrl: './rating-slider.component.html',
  standalone: true,
  styleUrl: './rating-slider.component.scss'
})
export class RatingSliderComponent {

  @Output()
  ratingChange = new EventEmitter<number>();

  result: string = "";

  constructor(private cdRef: ChangeDetectorRef) {}

  updateRating(value: number) {
    this.ratingChange.emit(value);
    this.result = this.formatTextLabel(value);
    this.cdRef.detectChanges();
  }

  private formatTextLabel(value: number): string {
     if (value === 0) {
     return "not fulfilled"
     } else if (value === 25) {
     return "slightly fulfilled"
     } else if (value === 50) {
     return "partially fulfilled"
     } else if (value === 75) {
     return "mostly fulfilled"
     } else if (value === 100) {
     return "fully fulfilled"
     } else if (value === 125) {
     return "overly fulfilled"
     } else {
     return "not rated"
     }
  }

  formatLabel(value: number): string {
    return `${value} %`;
  }

}
