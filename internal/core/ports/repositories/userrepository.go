package repositories

import (
<<<<<<< HEAD
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/domain"
=======
	"github.com/sugandhasaxena19/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/domain"
>>>>>>> 27c2bab4ae9973b95478eedee3812fd6c4b17ae0
)

type UserRepository interface {
	RegisterUser(domain.User) (domain.User, *error.AppError)
}
