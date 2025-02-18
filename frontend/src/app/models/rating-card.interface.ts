import {RatingCardDto} from './rating-card-dto.interface';


export interface RatingCard extends RatingCardDto {
  rating: number | undefined;
  response: string | undefined;
}
