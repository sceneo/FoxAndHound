import { Component, Input, OnChanges } from '@angular/core';

@Component({
  selector: 'app-rating-graphic',
  standalone: true,
  templateUrl: './rating-graphic.component.html',
  styleUrl: './rating-graphic.component.scss'
})
export class RatingGraphicComponent implements OnChanges {

  @Input()
  currentRating: number = 1.0;

  @Input()
  comparisonRating: number = 1.0;

  maxRating: number = 125;

  ratingPercentage: number = 0;
  comparisonPercentage: number = 0;

  ngOnChanges() {
    this.updatePercentages();
  }

  private updatePercentages() {
    this.ratingPercentage = Math.round((this.currentRating / this.maxRating) * 100);  // Round to nearest integer
    this.comparisonPercentage = Math.round((this.comparisonRating / this.maxRating) * 100);  // Round to nearest integer
  }

  getGradientColor(): string {
    return this.getColorForRating();
  }

  private getColorForRating(): string {
    const difference = this.ratingPercentage - this.comparisonPercentage;

    if (difference <= -75) {
      return '#ff0000'; // Red
    } else if (difference <= -50) {
      return '#ff6600'; // Orange
    } else if (difference <= -25) {
      return '#ffcc00'; // Yellow
    } else if (difference === 0) {
      return '#66ff00'; // Green
    } else {
      return '#0000ff'; // Blue
    }
  }

  getAverageRating(): number {
    return Math.floor(this.comparisonRating);
  }
}
