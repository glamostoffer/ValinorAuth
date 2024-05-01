package repository

import "github.com/glamostoffer/ValinorAuth/utils/pg_connector"

type PgClientRepository struct {
	db   *pg_connector.Connector
	user UserRepository
}

func New(db *pg_connector.Connector) *PgClientRepository {
	pg := &PgClientRepository{
		db: db,
	}

	pg.user = newUserRepo(pg)

	return pg
}
