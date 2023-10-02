package review

import (
	"context"
	"myapp/pkg/aws"
	"myapp/pkg/db/review"
)

type reviewservice struct {
	reviewDao review.Dao
	botocore  *aws.Boto3Wrapper
}

type ReviewService interface {
	GetReview(ctx context.Context, ReviewID uint) (review.Review, error)
}

func NewReviewService(r review.Dao, b *aws.Boto3Wrapper) ReviewService {
	return &reviewservice{
		reviewDao: r,
		botocore:  b,
	}
}

func (s *reviewservice) GetReview(ctx context.Context, ReviewID uint) (review.Review, error) {
	// Use the reviewDao to fetch the review by ID from the database
	review, err := s.reviewDao.GetReviewByID(ReviewID)
	if err != nil {
		return review, err // Return an empty review and the error
	}
	return review, nil
}

func (s *reviewservice) ListReviews(ctx context.Context, reviewIDs []uint) ([]review.Review, error) {
	reviews, err := s.reviewDao.ListReviews(reviewIDs)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (s *reviewservice) UpdateReview(ctx context.Context, reviewID uint, data map[string]interface{}) error {
	err := s.reviewDao.UpdateReview(reviewID, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *reviewservice) DeleteReview(ctx context.Context, reviewID uint) error {
	err := s.reviewDao.DeleteReview(reviewID)
	if err != nil {
		return err
	}
	return nil
}
