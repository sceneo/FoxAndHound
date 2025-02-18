import {Category} from './category.enum';
import {RatingCard} from './rating-card.interface';


export interface CategoryArrangement {
  name: string;
  category: Category;
  ratingCards?: RatingCard[];
}
