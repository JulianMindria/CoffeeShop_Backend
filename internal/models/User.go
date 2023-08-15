package models

import "time"

type User struct {
	Id_user    string     `db:"id_user" form:"id_user" json:"id_user,omitempty" valid:"-" `
	Username   string     `db:"username" form:"username" json:"username" valid:"-" `
	Email      string     `db:"email" form:"email" json:"email,omitempty" valid:"-" `
	Pass       string     `db:"pass" form:"pass" json:"pass,omitempty" valid:"stringlength(8|20)~create password from 8-20 character"`
	Phone      string     `db:"phone" form:"phone" json:"phone,omitempty" valid:"-" `
	Role       string     `db:"roles" form:"roles" json:"roles,omitempty" valid:"-" `
	Image_file string     `db:"image_file" json:"image_file,omitempty" valid:"-" `
	Created_at *time.Time `db:"created_at" json:"created_at" valid:"-" `
	Updated_at *time.Time `db:"updated_at" json:"updated_at" valid:"-" `
}
