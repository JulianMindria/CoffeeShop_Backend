package repositories

import (
	"julianmindria/backendCoffee/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoFav struct {
	*sqlx.DB
}

func NewFav(db *sqlx.DB) *RepoFav {
	return &RepoFav{db}
}

func (r RepoFav) Create_favproduct(data *models.Fav_Product) (string, error) {
	queryproduct := `INSERT INTO favorite_products(
		id_user,
		id_product  
		) 
		VALUES(
			:id_user, 
			:id_product 
		)`

	_, err := r.NamedExec(queryproduct, data)
	if err != nil {
		return "", err
	}
	return "1 favorite product Added", nil

}
