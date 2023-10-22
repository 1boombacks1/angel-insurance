package models

type CarPhoto struct {
	Id             int    `db:"id"`
	CarInsuranceId int    `db:"car_insurance_id"`
	TypePhoto      string `db:"type_photo"`
	LinkToPhoto    string `db:"link_to_photo"`
}
