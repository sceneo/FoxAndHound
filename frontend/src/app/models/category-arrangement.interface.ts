import { ModelsRatingCard } from '../api';
import {Category} from './category.enum';


export interface CategoryArrangement {
  name: string;
  category: Category;
  ratingCards?: ModelsRatingCard[];
}
