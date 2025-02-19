import {Component, EventEmitter, Output} from '@angular/core';
import {MatSlider, MatSliderThumb} from '@angular/material/slider';
import {elementAt} from 'rxjs';

@Component({
  selector: 'app-rating-slider',
  imports: [
    MatSlider,
    MatSliderThumb
  ],
  templateUrl: './rating-slider.component.html',
  standalone: true,
  styleUrl: './rating-slider.component.scss'
})
export class RatingSliderComponent {

  @Output()
  ratingChange = new EventEmitter<number>();

  updateRating(value: number) {
    this.ratingChange.emit(value);
  }

  formatLabel(value: number): string {
    return `${value} %`;
    // TODO check if we prefer that:
    /**
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
      **/
  }

}
