import {Injectable} from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ContentService {
  static getMainHeader(): string {
    return "Senior Suggestion Service"
  }
}
