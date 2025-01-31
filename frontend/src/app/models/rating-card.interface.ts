import {Category} from './category.enum';

export interface RatingCard {
  id: string;
  question: string;
  category: Category;
  orderId: number;
}
