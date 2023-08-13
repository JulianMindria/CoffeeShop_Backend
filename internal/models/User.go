package models

import "time"

type User struct {
	Id_user     string     `db:"id_user" form:"id_user" json:"id_user,omitempty" `
	Username    string     `db:"username" form:"username" json:"username" `
	Email       string     `db:"email" form:"email" json:"email,omitempty" `
	Pass        string     `db:"pass" form:"pass" json:"pass,omitempty" `
	Phone       string     `db:"phone" form:"phone" json:"phone,omitempty" `
	Role        string     `db:"roles" form:"roles" json:"roles,omitempty" `
	User_image  string     `db:"user_image" json:"user_image,omitempty" `
	Fav_Product string     `db:"fav_product" form:"fav_product" json:"fav_product" `
	Created_at  *time.Time `db:"created_at" json:"created_at" `
	Updated_at  *time.Time `db:"updated_at" json:"updated_at" `
}
