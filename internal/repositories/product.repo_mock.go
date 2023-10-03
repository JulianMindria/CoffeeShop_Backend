package repositories

import (
	"julianmindria/backendCoffee/config"
	"julianmindria/backendCoffee/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoProductMock struct {
	mock.Mock
}

func (rpd *RepoProductMock) Createproduct(data *models.Product) (*config.Result, error) {
	args := rpd.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (rpd *RepoProductMock) Updateproduct(data *models.Product) (*config.Result, error) {
	args := rpd.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}
func (rpd *RepoProductMock) Deleteproduct(data *models.Product) (*config.Result, error) {
	args := rpd.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}
func (rpd *RepoProductMock) Get_products(data *models.Product, search string, limit int, page int, orderby string) ([]models.Product, *Pagination, error) {
	args := rpd.Mock.Called(data)
	return args.Get(0).([]models.Product), args.Get(0).(*Pagination), args.Error(0)
}
