package models

import "time"

type Order_Product struct {
	Order_id   string     `db:"order_id" form:"order_id" json:"order_id" `
	Id_user    string     `db:"id_user" form:"id_user" json:"id_user" `
	Id_product string     `db:"id_product" form:"id_product" json:"id_product" `
	Status     string     `db:"status" form:"status" json:"status"`
	Created_at *time.Time `db:"created_at" json:"created_at" `
	Updated_at *time.Time `db:"updated_at" json:"updated_at" `
}
