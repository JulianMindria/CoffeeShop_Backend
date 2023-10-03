package repositories

import (
	"julianmindria/backendCoffee/config"
	"julianmindria/backendCoffee/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoUserMock struct {
	mock.Mock
}

func (rp *RepoUserMock) CreateUser(data *models.User) (*config.Result, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (rp *RepoUserMock) UpdateUser(data *models.User) (*config.Result, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}
func (rp *RepoUserMock) DeleteUser(data *models.User) (*config.Result, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}
func (rp *RepoUserMock) Get_Users(data *models.User, search string) ([]models.User, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).([]models.User), args.Error(1)
}
func (rp *RepoUserMock) GetAuthUser(user string) (*models.User, error) {
	args := rp.Mock.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}
