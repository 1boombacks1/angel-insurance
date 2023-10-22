package models

import "time"

type Client struct {
	Id         int       `db:"id"`
	LastName   string    `db:"last_name"`
	FirstName  string    `db:"first_name"`
	Patronymic string    `db:"patronymic"`
	Phone      string    `db:"phone"`
	LinkToChat string    `db:"link_to_chat"`
	Login      string    `db:"login"`
	Password   string    `db:"password"`
	CreatedAt  time.Time `db:"created_at"`
}
