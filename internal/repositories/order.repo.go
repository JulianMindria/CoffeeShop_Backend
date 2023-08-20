package repositories

import (
	"julianmindria/backendCoffee/config"
	"julianmindria/backendCoffee/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoOrder struct {
	*sqlx.DB
}

func NewFav(db *sqlx.DB) *RepoOrder {
	return &RepoOrder{db}
}

func (r RepoOrder) Add_Order_product(data *models.Order_Product) (*config.Result, error) {
	queryproduct := `INSERT INTO order_products(
		id_user,
		id_product 
		) 
		VALUES(
			:id_user, 
			:id_product 
		)`

	_, err := r.NamedExec(queryproduct, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 Order created"}, nil
}

func (r RepoOrder) Delete_Order_product(data *models.Order_Product) (*config.Result, error) {
	queryproduct := `DELETE FROM public.order_products WHERE order_id=:order_id`

	_, err := r.NamedExec(queryproduct, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 Order deleted"}, nil
}
