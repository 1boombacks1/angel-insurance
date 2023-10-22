package models

import "time"

const (
	CodeCreated = iota
	CodeInProcessing
	CodeRejected
	CodeAccepted
)

type InsuranceApplications struct {
	Id        int       `db:"id"`
	ClientID  int       `db:"client_id"`
	InsurerID int       `db:"insurer_id"`
	Status    int       `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}
