package repositories

import (
	"absensi-app/backend/configs"
	"absensi-app/backend/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	db configs.Config
}

func (r *UserRepository) GetAll(page, limit int, name, role, email string) ([]models.User, int, error) {
	var users []models.User
	var total int64

	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	query := r.db.GetDb().Preload("Activity").Model(&models.User{}).Select("id", "name", "image", "first_login", "role", "status", "created_at", "updated_at", "Email")

	query = query.Where("role NOT LIKE ?", "ROLE_SUPER_ADMIN")

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	if role != "" {
		query = query.Where("role LIKE ?", role)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Limit(limit).Offset(offset).Order("id asc").Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (r *UserRepository) GetById(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.GetDb().First(&user, id).Error
	return &user, err
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.GetDb().Create(user).Error
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.GetDb().Save(&user).Error
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	return r.db.GetDb().Delete(&models.User{}, id).Error
}

func (r *UserRepository) Count() (int64, error) {
	var count int64
	err := r.db.GetDb().Model(&models.User{}).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}
