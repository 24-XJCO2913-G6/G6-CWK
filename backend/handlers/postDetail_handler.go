package handlers

import (
	. "main/backend/models"
)

func GetBlog(Id int64) (Blog, error) {
	var blog Blog
	err := Db.Where("Id = ?", Id).Find(&blog)
	if err != nil {
		return Blog{}, err
	}
	return blog, nil
}

func GetReviews(Id int64) ([]ReviewInBlog, error) {
	var reviews []Review
	var reviews_in_blog []ReviewInBlog

	err := Db.Where("Bid = ?", Id).Find(&reviews)
	if err != nil {
		return []ReviewInBlog{}, err
	}
	for _, review := range reviews {
		var reviewers []User
		err := Db.Where("Uid = ?", review.Uid).Find(&reviewers)
		if err != nil {
			return []ReviewInBlog{}, err
		}
		reviewInBlog := ReviewInBlog{Time: review.Time, Content: review.Content, Reviewer: reviewers[0].Name, Reviewer_photo: reviewers[0].ProfilePhot}
		reviews_in_blog = append(reviews_in_blog, reviewInBlog)
	}
	return reviews_in_blog, nil
}
