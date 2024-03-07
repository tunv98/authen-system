package database

import (
	"authen-system/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	Create(user *models.User) error
	Find(queries map[string]string) (models.User, error)
	UpdateLatestLogin(userID uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) Find(queries map[string]string) (models.User, error) {
	var user models.User
	r.db = buildGormQuery(r.db, queries)
	if err := r.db.First(&user).Error; err != nil {
		return user, errors.Wrapf(err, "failed to find user")
	}
	return user, nil
}

func (r *userRepo) UpdateLatestLogin(userID uint) error {
	now := time.Now()
	return r.db.
		Model(&models.User{}).
		Where("id = ?", userID).
		Update("latest_login", &now).
		Error
}
