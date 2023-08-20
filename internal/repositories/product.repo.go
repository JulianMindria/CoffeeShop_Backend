package repositories

import (
	"errors"
	"fmt"
	"julianmindria/backendCoffee/config"
	"julianmindria/backendCoffee/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

type Pagination struct {
	Next          int
	Previous      int
	RecordPerPage int
	CurrentPage   int
	TotalPage     int
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r RepoProduct) Createproduct(data *models.Product) (*config.Result, error) {
	queryproduct := `INSERT INTO products(
		product_name,  
		description, 
		stock,         
		price,         
		image_file,
		categories,
		isfavorite
		) 
		VALUES(
			:product_name, 
			:description, 
			:stock, 
			:price,
			:image_file,
			:categories,
			:isfavorite
		)`

	_, err := r.NamedExec(queryproduct, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 product user created"}, nil

}

func (r RepoProduct) Updateproduct(data *models.Product) (*config.Result, error) {
	queryproduct := `UPDATE public.products
	SET product_name=:product_name, description=:description, stock=:stock, price=:price, image_file=:image_file, categories=:categories, isfavorite=:isfavorite
	WHERE id_product=:id_product
	`
	_, err := r.NamedExec(queryproduct, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 product user updated"}, nil

}

func (r RepoProduct) Deleteproduct(data *models.Product) (*config.Result, error) {
	queryproduct := `DELETE FROM public.products WHERE id_product=:id_product`
	_, err := r.NamedExec(queryproduct, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 product user deleted"}, nil

}

func (r RepoProduct) Get_products(data *models.Product, search string, limit int, page int, orderby string) ([]models.Product, *Pagination, error) {
	products_data := []models.Product{}

	var (
		pgnt        = &Pagination{}
		recordcount int
	)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 5
	}

	offset := limit * (page - 1)

	sqltable := fmt.Sprintf("SELECT count(id_product) FROM products")

	r.QueryRow(sqltable).Scan(&recordcount)

	total := (recordcount / limit)

	remainder := (recordcount % limit)
	if remainder == 0 {
		pgnt.TotalPage = total
	} else {
		pgnt.TotalPage = total + 1
	}

	pgnt.CurrentPage = page
	pgnt.RecordPerPage = limit

	if page <= 0 {
		pgnt.Next = page + 1
	} else if page < pgnt.TotalPage {
		pgnt.Previous = page - 1
		pgnt.Next = page + 1
	} else if page == pgnt.TotalPage {
		pgnt.Previous = page - 1
		pgnt.Next = 0
	}

	if search != "" {
		r.Select(&products_data, `select * from products WHERE categories=$1 LIMIT $2 OFFSET $3 `, search, limit, offset)
	} else if orderby != "" {
		r.Select(&products_data, `select * from products order by `+orderby+` asc LIMIT $1 OFFSET $2`, limit, offset)
	} else if orderby != "" && search != "" {
		r.Select(&products_data, `select * from products order by `+orderby+` asc WHERE categories=$1 LIMIT $2 OFFSET $3 `, search, limit, offset)
	} else {
		r.Select(&products_data, `select * from products LIMIT $1 OFFSET $2 `, limit, offset)
	}
	if len(products_data) == 0 {
		return nil, nil, errors.New("data not found.")
	}
	return products_data, pgnt, nil
}
