import { Component, forwardRef } from '@angular/core';
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';

@Component({
  selector: 'app-rating-slider',
  standalone: true,
  imports: [
    MatSliderModule,
    MatFormFieldModule
  ],
  templateUrl: './rating-slider.component.html',
  styleUrl: './rating-slider.component.scss',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => RatingSliderComponent),
      multi: true
    }
  ]
})
export class RatingSliderComponent implements ControlValueAccessor {
  selectedValue: number = 0;
  result: string = "";
  isDisabled = false;

  private onChange = (value: number) => {};
  private onTouched = () => {};

  updateRating(value: number) {
    this.selectedValue = value;
    this.result = this.formatTextLabel(value);
    this.onChange(value);
    this.onTouched();
  }

  writeValue(value: number): void {
    this.selectedValue = value;
    this.result = this.formatTextLabel(value);
  }

  registerOnChange(fn: (value: number) => void): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: () => void): void {
    this.onTouched = fn;
  }

  formatTextLabel(value: number): string {
    switch (value) {
      case 0: return "not fulfilled";
      case 25: return "slightly fulfilled";
      case 50: return "partially fulfilled";
      case 75: return "mostly fulfilled";
      case 100: return "fully fulfilled";
      case 125: return "overly fulfilled";
      default: return "not rated";
    }
  }

  formatLabel(value: number): string {
    return `${value} %`;
  }

  setDisabledState(isDisabled: boolean): void {
      this.isDisabled = isDisabled;
  }
}
