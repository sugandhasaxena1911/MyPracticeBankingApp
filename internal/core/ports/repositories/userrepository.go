package repositories

import (
	"github.com/sugandhasaxena19/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/domain"
)

type UserRepository interface {
	RegisterUser(domain.User) (domain.User, *error.AppError)
}
