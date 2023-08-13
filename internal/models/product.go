package models

import "time"

type Product struct {
	Id_product    string     `db:"id_product" form:"id_product" json:"id_product" `
	Product_name  string     `db:"product_name" form:"product_name" json:"product_name" `
	Description   string     `db:"description" form:"description" json:"description" `
	Stock         int64      `db:"stock" form:"stock" json:"stock"`
	Price         int64      `db:"price" form:"price" json:"price"`
	Product_image string     `db:"product_image" form:"product_image" json:"product_image" `
	Categories    string     `db:"categories" form:"categories" json:"categories" `
	Created_at    *time.Time `db:"created_at" json:"created_at" `
	Updated_at    *time.Time `db:"updated_at" json:"updated_at" `
}
