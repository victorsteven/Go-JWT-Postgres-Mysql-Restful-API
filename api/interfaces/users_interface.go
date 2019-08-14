package interfaces

import "github.com/victorsteven/fullstack/api/models"

type UserInterface interface {
	SaveUser(models.User) (models.User, error)
	FindAllUsers() ([]models.User, error)
	FindUserByID(uint32) (models.User, error)
	UpdateAUser(uint32, models.User) (models.User, error)
	DeleteAUser(uint32) (int64, error)
}
