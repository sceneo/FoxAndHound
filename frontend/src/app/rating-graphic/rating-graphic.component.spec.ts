import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RatingGraphicComponent } from './rating-graphic.component';

describe('RatingSliderComponent', () => {
  let component: RatingGraphicComponent;
  let fixture: ComponentFixture<RatingGraphicComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RatingGraphicComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RatingGraphicComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
