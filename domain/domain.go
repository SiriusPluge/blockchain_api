package domain

import "blockchain_api/internal/postgres"

type Domain struct {
	DB postgres.PostgresDB
}

func NewDomain(DB postgres.PostgresDB) *Domain {
	return &Domain{DB: DB}
}
