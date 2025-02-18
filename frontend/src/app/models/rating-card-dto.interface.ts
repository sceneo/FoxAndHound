import {Category} from './category.enum';

export interface RatingCardDto {
  id: string;
  question: string;
  category: Category;
  orderId: number;
}
