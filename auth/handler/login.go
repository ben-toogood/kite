package handler

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/auth/model"
	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/users"
	"github.com/dgrijalva/jwt-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login sends an access token to the users email
func (a *Auth) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	u, err := a.Users.GetByEmail(ctx, &users.GetByEmailRequest{Email: req.Email})
	if status.Code(err) == codes.NotFound {
		return &auth.LoginResponse{}, nil
	}

	if err != nil {
		return nil, err
	}

	// calculate the expiry
	accessTokenExpiry := time.Now().Add(accessTokenTTL)
	refreshTokenExpiry := time.Now().Add(refreshTokenTTL)

	// generate the tokens
	at := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: accessTokenExpiry.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   u.User.Id,
	})
	rt := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: refreshTokenExpiry.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   u.User.Id,
	})

	// sign the tokens
	accessToken, err := at.SignedString(a.PrivateKey)
	if err != nil {
		return nil, err
	}
	refreshToken, err := rt.SignedString(a.PrivateKey)
	if err != nil {
		return nil, err
	}

	// write the token to the store
	tok := &model.Token{
		UserID:             u.User.Id,
		AccessToken:        accessToken,
		AccessTokenExpiry:  accessTokenExpiry,
		RefreshToken:       refreshToken,
		RefreshTokenExpiry: refreshTokenExpiry,
	}
	if err := a.DB.Create(tok).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	// construct the url
	vals := make(url.Values)
	vals.Add("code", tok.RefreshToken)
	url := fmt.Sprintf("%v/login?%v", os.Getenv("KITE_WEB_URL"), vals.Encode())

	// send the email
	from := mail.NewEmail("Kite", "support@m3o.com")
	subject := "Login with Kite"
	to := mail.NewEmail("Example User", req.Email)
	plainTextContent := fmt.Sprintf("Copy and paste %v into your browser", url)
	htmlContent := fmt.Sprintf("<a href=\"%v\">Click here</a>", url)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	if _, err := a.Sendgrid.Send(message); err != nil {
		return nil, err
	}

	return &auth.LoginResponse{}, nil
}
