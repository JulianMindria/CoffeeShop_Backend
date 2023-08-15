package models

import "time"

type Product struct {
	Id_product   string     `db:"id_product" form:"id_product" json:"id_product" `
	Product_name string     `db:"product_name" form:"product_name" json:"product_name" `
	Description  string     `db:"description" form:"description" json:"description" `
	Stock        int64      `db:"stock" form:"stock" json:"stock"`
	Price        int64      `db:"price" form:"price" json:"price"`
	Image_file   string     `db:"image_file" json:"image_file,omitempty" valid:"-" `
	Categories   string     `db:"categories" form:"categories" json:"categories" `
	IsFavorite   string     `db:"isfavorite" form:"isfavorite" json:"isfavorite" `
	Created_at   *time.Time `db:"created_at" json:"created_at" `
	Updated_at   *time.Time `db:"updated_at" json:"updated_at" `
}
