import { TestBed } from '@angular/core/testing';
import { ManagementDashboardComponent } from './management-dashboard.component';

describe('AppComponent', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ManagementDashboardComponent],
    }).compileComponents();
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(ManagementDashboardComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  it(`should have the 'frontend' title`, () => {
    const fixture = TestBed.createComponent(ManagementDashboardComponent);
    const app = fixture.componentInstance;
    expect(app.title).toEqual('frontend');
  });

  it('should render title', () => {
    const fixture = TestBed.createComponent(ManagementDashboardComponent);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h1')?.textContent).toContain('Hello, frontend');
  });
});
