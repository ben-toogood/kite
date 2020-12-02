package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/auth/model"
	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/dgrijalva/jwt-go"
	"github.com/lileio/lile/v2/protocopy"
	"gorm.io/gorm"
)

// Refresh an access token
func (a *Auth) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// lookup the refresh token in the store
	q := a.DB.Where("refresh_token = ? AND refresh_token_expiry >= ?", req.RefreshToken, time.Now())
	var old model.Token
	if err := q.First(&old).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Refresh token not found")
	} else if err != nil {
		return nil, err
	}

	// calculate the expiry
	accessTokenExpiry := time.Now().Add(accessTokenTTL)
	refreshTokenExpiry := time.Now().Add(refreshTokenTTL)

	// generate the new tokens
	at := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: accessTokenExpiry.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   old.UserID,
	})
	rt := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: refreshTokenExpiry.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   old.UserID,
	})

	// sign the tokens
	accessToken, err := at.SignedString(a.JWTPrivateKey)
	if err != nil {
		return nil, err
	}
	refreshToken, err := rt.SignedString(a.JWTPrivateKey)
	if err != nil {
		return nil, err
	}

	// write the token to the store
	tok := &model.Token{
		UserID:             old.UserID,
		AccessToken:        accessToken,
		AccessTokenExpiry:  accessTokenExpiry,
		RefreshToken:       refreshToken,
		RefreshTokenExpiry: refreshTokenExpiry,
	}
	if err := a.DB.Create(tok).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	// serialize the token
	rsp := auth.RefreshResponse{Token: &auth.Token{}}
	if err := protocopy.ToProto(tok, rsp.Token); err != nil {
		return nil, err
	}
	return &rsp, nil
}
