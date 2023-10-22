package models

type CarInsurance struct {
	Id            int    `db:"id"`
	ApplicationId int    `db:"application_id"`
	Description   string `db:"description"`
}
