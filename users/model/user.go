package model

import (
	"context"
	"time"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/users"
	"github.com/lileio/lile/v2/protocopy"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Serialize() (*users.User, error) {
	pbu := users.User{}
	return &pbu, protocopy.ToProto(u, &pbu)
}

func (u *User) BeforeCreate(scope *gorm.DB) error {
	u.ID = "usr_" + ksuid.New().String()
	return nil
}

func (u *User) AfterCreate(scope *gorm.DB) error {
	return users.PublishUserCreated(scope.Statement.Context, u)
}

func Create(ctx context.Context, u *User) error {
	db, err := database.GetDB(ctx)
	if err != nil {
		return err
	}

	return database.TranslateError(db.Create(u).Error)
}

func Get(ctx context.Context, ids []string) (map[string]*User, error) {
	db, err := database.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	var usrs []*User
	if err := db.Where("id IN (?)", ids).Find(&usrs).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	umap := map[string]*User{}
	for _, u := range usrs {
		umap[u.ID] = u
	}

	return umap, nil
}

func GetByEmail(ctx context.Context, email string) (*User, error) {
	db, err := database.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	var u User
	if err := db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, database.TranslateError(err)
	}

	return &u, nil
}
