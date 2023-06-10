package ratingbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product_rating/ratingmodel"
	"context"
)

type UpdateRatingStore interface {
	UpdateRating(context context.Context, id int, data *ratingmodel.RatingUpdate) error
	FindDataWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ratingmodel.Rating, error)
}

type updateRatingBusiness struct {
	store UpdateRatingStore
}

func NewUpdateRatingBusiness(store UpdateRatingStore) *updateRatingBusiness {
	return &updateRatingBusiness{store: store}
}

func (business *updateRatingBusiness) UpdateRating(context context.Context, userId, id int, data *ratingmodel.RatingUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	result, err := business.store.FindDataWithCondition(context, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return err
	}

	if result.UserID != userId {
		return common.ErrNoPermission(nil)
	}

	if result == nil {
		return common.ErrEntityNotFound(ratingmodel.EntityName, nil)
	}

	if result.Status == 0 {
		return common.ErrEntityDeleted(ratingmodel.EntityName, nil)
	}

	if err := business.store.UpdateRating(context, id, data); err != nil {
		return err
	}
	return nil
}
