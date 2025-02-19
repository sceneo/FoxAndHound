import {Component, EventEmitter, Output} from '@angular/core';
import {MatSlider, MatSliderThumb} from '@angular/material/slider';

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


}
