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

func (s *UserService) CreateUser(createUser *models.CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := &models.User{
		Name:     createUser.Name,
		Email:    createUser.Email,
		Password: string(hashedPassword),
	}

	count, _ := s.userRepo.Count()
	if count == 0 {
		user.Role = models.RoleAdmin
	}

	return s.userRepo.Create(user)
}

func (s *UserService) UpdateUser(userId string, userUpdate *models.UpdateUserRequest) error {
	user, err := s.userRepo.GetById(uuid.MustParse(userId))

	if err != nil {
		return err
	}

	if userUpdate.Name != "" {
		user.Name = userUpdate.Name
	}

	if userUpdate.Role != "" {
		user.Role = models.Role(userUpdate.Role)
	}

	if userUpdate.Status != "" {
		user.Status = models.UserStatus(userUpdate.Status)
	}

	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(userId string) error {
	return s.userRepo.Delete(uuid.MustParse(userId))
}
