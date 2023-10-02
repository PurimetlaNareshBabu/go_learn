package review

import (
	"errors"
	"myapp/pkg/db"

	"gorm.io/gorm"
)

type reviewdao struct {
	db *gorm.DB
}

type Dao interface {
	Createreview(newreview *Review) (uint, error)
	UpdateReview(review_id uint, data map[string]interface{}) error
	GetReviewByID(review_id uint) (Review, error)
	ListReviews(reviewIDs []uint) ([]Review, error)
	DeleteReview(review_id uint) error
}

func NewReviewDao(dbClient *db.Client) *reviewdao {
	return &reviewdao{
		db: dbClient.DB,
	}
}

func (d *reviewdao) CreateReview(newreview *Review) (uint, error) {
	err := d.db.Create(newreview)
	if err != nil {
		return 0, errors.New("failed to create review: " + err.Error.Error())
	}
	return newreview.ID, nil
}

func (d *reviewdao) UpdateReview(review_id uint, data map[string]interface{}) error {
	err := d.db.Model(&Review{}).Where("id = ?", review_id).Updates(data)
	if err != nil {
		return errors.New("failed to update review: " + err.Error.Error())
	}
	return nil
}

func (d *reviewdao) GetReviewByID(review_id uint) (Review, error) {
	var review Review
	err := d.db.First(&review, review_id).Error
	if err != nil {
		return Review{}, err
	}
	return review, nil
}

func (d *reviewdao) ListReviews(reviewIDs []uint) ([]Review, error) {
	var reviews []Review
	err := d.db.Where("id IN (?)", reviewIDs).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (d *reviewdao) Deletereview(review_id uint) error {
	err := d.db.Delete(&Review{}, review_id).Error
	if err != nil {
		return err
	}
	return nil
}
