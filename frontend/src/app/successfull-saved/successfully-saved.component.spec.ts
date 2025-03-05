import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SuccessfullySavedComponent } from './successfully-saved.component';

describe('SuccessfullComponent', () => {
  let component: SuccessfullySavedComponent;
  let fixture: ComponentFixture<SuccessfullySavedComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SuccessfullySavedComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SuccessfullySavedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
