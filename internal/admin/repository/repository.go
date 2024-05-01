package repository

import "github.com/glamostoffer/ValinorAuth/utils/pg_connector"

type PgClientRepository struct {
	db    *pg_connector.Connector
	admin AdminRepository
}

func New(db *pg_connector.Connector) *PgClientRepository {
	pg := &PgClientRepository{
		db: db,
	}

	pg.admin = newAdminRepo(pg)

	return pg
}
