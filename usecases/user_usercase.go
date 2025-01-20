package usecases

import "userapi/domain"

type UserUsecase struct {
	UserRepository domain.UserRepository
}


func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.UserRepository.CreateUser(user)
}

func (u *UserUsecase) GetUserByID(id int) (*domain.User, error) {
	return u.UserRepository.GetUserByID(id)
}

func (u *UserUsecase) UpdateUser(user *domain.User) error {
	return u.UserRepository.UpdateUser(user)
}

func (u *UserUsecase) DeleteUser(id int) error {
	return u.UserRepository.DeleteUser(id)
}
