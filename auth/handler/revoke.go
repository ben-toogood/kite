package handler

import (
	"context"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/auth/model"
	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
)

// Revoke all tokens for a user
func (a *Auth) Revoke(ctx context.Context, req *auth.RevokeRequest) (*auth.RevokeResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// delete the tokens
	if err := a.DB.Where(&model.Token{UserID: req.UserId}).Delete(&model.Token{}).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	return &auth.RevokeResponse{}, nil
}
