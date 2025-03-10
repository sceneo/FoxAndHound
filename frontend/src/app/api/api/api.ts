export * from './headData.service';
import { HeadDataService } from './headData.service';
export * from './ratingCandidate.service';
import { RatingCandidateService } from './ratingCandidate.service';
export * from './ratingCard.service';
import { RatingCardService } from './ratingCard.service';
export * from './ratingEmployer.service';
import { RatingEmployerService } from './ratingEmployer.service';
export const APIS = [HeadDataService, RatingCandidateService, RatingCardService, RatingEmployerService];
