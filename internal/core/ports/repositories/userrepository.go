package repositories

import (
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/domain"
)

type UserRepository interface {
	RegisterUser(domain.User) (domain.User, *error.AppError)
}
