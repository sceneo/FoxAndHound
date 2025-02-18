import {Injectable} from '@angular/core';
import {CandidateDto} from './model/candidateDto';
import {HttpClient} from '@angular/common/http';
import {RatingCardDto} from '../models/rating-card-dto.interface';
import {Observable} from 'rxjs';


@Injectable()
export class CandidateApiService {

  // TODO: make url configurable
  private apiUrl = 'http://localhost:8080/api/candidates';

  constructor(private http: HttpClient) {
  }


  getCandidates(): Observable<CandidateDto[]> {
    return this.http.get<CandidateDto[]>(this.apiUrl);
  }

}
