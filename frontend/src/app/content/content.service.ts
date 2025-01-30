import {Injectable} from '@angular/core';
import {CategoryArrangement} from '../models/category-arrangement.interface';
import {Category} from '../models/category.enum';

@Injectable({
  providedIn: 'root',
})
export class ContentService {
  static getMainHeader(): string {
    return "Senior Suggestion Service"
  }

  static getCategoryArrangement(): CategoryArrangement[] {
    return [
      { name: "Job / Customer Performance", category: Category.PERFORMANCE },
      { name: "Professional Skillset & Learning", category: Category.TECHNICAL_SKILLSET },
      { name: "Technical Predispositions", category: Category.TECHNICAL_PREDISPOSITIONS },
      { name: "Sales Support", category: Category.SALES },
      { name: "HR/Recruiting", category: Category.RECRUITING },
      { name: "Teamwork & Social Skills", category: Category.TEAMWORK },
      { name: "Mentoring & Coaching", category: Category.COACHING },
      { name: "PRODYNA Insights", category: Category.PRODYNA_INSIGHTS },
      { name: "Overall Performance", category: Category.OVERALL }
  ];
  }
}
