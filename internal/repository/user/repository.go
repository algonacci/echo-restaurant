package user

import "github.com/algonacci/echo-restaurant/internal/model"

type Repository interface {
	RegisterUser(userData model.User) (model.User, error)
	CheckRegistered(username string) (bool, error)
	GenerateUserHash(password string) (hash string, err error)
	VerifyLogin(username, password string, userData model.User) (bool, error)
	GetUserData(username string) (model.User, error)
	CreateUserSession(userId string) (model.UserSession, error)
	CheckSession(data model.UserSession) (userID string, err error)
}
