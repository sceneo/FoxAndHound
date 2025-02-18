import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {RatingCard} from '../models/rating-card.interface';
import {HttpClient} from '@angular/common/http';


@Injectable()
export class RatingCardApiService {

  // TODO: make apiUrl configurable
  private apiUrl = 'http://localhost:8080/api/rating-cards';

  constructor(private http: HttpClient) {}

  getRatingCards(): Observable<RatingCard[]> {
    return this.http.get<RatingCard[]>(this.apiUrl);
  }

}

