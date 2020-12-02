package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/auth/model"
	"gorm.io/gorm"
)

// Inspect an access token
func (a *Auth) Inspect(ctx context.Context, req *auth.InspectRequest) (*auth.InspectResponse, error) {
	// lookup the token in the store
	var q *gorm.DB
	if len(req.AccessToken) > 0 {
		q = a.DB.Where("access_token = ? AND access_token_expiry >= ?", req.AccessToken, time.Now())
	} else if len(req.RefreshToken) > 0 {
		q = a.DB.Where("refresh_token = ? AND refresh_token_expiry >= ?", req.RefreshToken, time.Now())
	} else {
		return nil, errors.New("Missing AccessToken or RefreshToken")
	}

	var tok model.Token
	if err := q.First(&tok).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Access token not found")
	} else if err != nil {
		return nil, err
	}

	// return the user id
	return &auth.InspectResponse{UserId: tok.UserID}, nil
}
