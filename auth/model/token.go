package model

import "time"

type Token struct {
	UserID             string `pb:"ignore=true"`
	AccessToken        string
	RefreshToken       string
	AccessTokenExpiry  time.Time
	RefreshTokenExpiry time.Time
}
