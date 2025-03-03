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

    let red, green, blue;

    // Below -60, always red
    if (difference <= -60) {
      red = 255;
      green = 0;
      blue = 0;
    }
    // Above +25, always blue
    else if (difference >= 25) {
      red = 0;
      green = 0;
      blue = 255;
    }
    // Smooth transition between -60 and +25
    else {
      // Adjusting transitions based on current average ratings
      const transitionRange = 25;  // Range where transitions happen

      // Red to Orange (for negative values)
      if (difference < -5) {
        red = 255;
        green = Math.min(255, Math.floor(255 + difference * (255 / transitionRange))); // Adjusted for better scaling
        blue = 0;
      }
      // Orange to Yellow to Green (around 0)
      else if (difference <= -5) {
        red = Math.min(255, Math.floor(255 - difference * (255 / transitionRange))); // Smooth orange to yellow to green
        green = 255;
        blue = 0;
      }
      // Green to Blue (for positive values)
      else {
        red = 0;
        green = Math.min(255, Math.floor(255 - (difference - 0) * (255 / transitionRange))); // Smooth green to blue transition
        blue = Math.min(255, Math.floor((difference - 0) * (255 / transitionRange))); // Smooth green to blue transition
      }
    }

    // Return the RGB color
    return `rgb(${red}, ${green}, ${blue})`;
  }




  getAverageRating(): number {
    return Math.floor(this.comparisonRating);
  }
}
