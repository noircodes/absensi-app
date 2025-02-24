package services

import (
	"absensi-app/backend/models"
	"absensi-app/backend/repositories"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func (s *UserService) GetAllUsers(page, limit int, name, role, email string) ([]models.User, int, error) {
	return s.userRepo.GetAll(page, limit, name, role, email)
}

func (s *UserService) GetUserById(id string) (*models.User, error) {
	return s.userRepo.GetById(uuid.MustParse(id))
}

func (s *UserService) CreateUser(userRegister *models.RegisterUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := &models.User{
		Name:     userRegister.Name,
		Email:    userRegister.Email,
		Password: string(hashedPassword),
	}

	count, _ := s.userRepo.Count()
	if count == 0 {
		user.Role = models.RoleAdmin
	}

	return s.userRepo.Create(user)
}
