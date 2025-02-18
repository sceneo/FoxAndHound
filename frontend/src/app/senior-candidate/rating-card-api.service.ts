import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {RatingCardDto} from '../models/rating-card-dto.interface';
import {HttpClient} from '@angular/common/http';


@Injectable()
export class RatingCardApiService {

  // TODO: make apiUrl configurable
  private apiUrl = 'http://localhost:8080/api/rating-cards';

  constructor(private http: HttpClient) {}

  getRatingCards(): Observable<RatingCardDto[]> {
    return this.http.get<RatingCardDto[]>(this.apiUrl);
  }

}

