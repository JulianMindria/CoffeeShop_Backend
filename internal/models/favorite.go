package models

import "time"

type Fav_Product struct {
	Id_user    string     `db:"id_user" form:"id_user" json:"id_user" `
	Id_product string     `db:"id_product" form:"id_product" json:"id_product" `
	Created_at *time.Time `db:"created_at" json:"created_at" `
	Updated_at *time.Time `db:"updated_at" json:"updated_at" `
}
