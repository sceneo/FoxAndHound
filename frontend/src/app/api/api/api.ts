export * from './ratingCandidate.service';
import { RatingCandidateService } from './ratingCandidate.service';
export * from './ratingCard.service';
import { RatingCardService } from './ratingCard.service';
export * from './ratingEmployer.service';
import { RatingEmployerService } from './ratingEmployer.service';
export const APIS = [RatingCandidateService, RatingCardService, RatingEmployerService];
