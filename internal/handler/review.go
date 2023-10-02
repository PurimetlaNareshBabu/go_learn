package handler

import (
	"myapp/internal/modules/review"
)

type ReviewHandler struct {
	service *review.ReviewService
}

func NewReviewHandler(s *review.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		service: s,
	}
}

// func (r *ReviewHandler)getReviewsHandler(c echo.Context) error {
// 	// Fetch reviews from the database
// 	var reviews []Review
// 	result := db.Find(&reviews)
// 	if result.Error != nil {
// 		return c.String(http.StatusInternalServerError, "Error fetching reviews")
// 	}

// 	return c.JSON(http.StatusOK, reviews)
// }
