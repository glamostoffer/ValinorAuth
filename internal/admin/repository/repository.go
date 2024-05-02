package repository

import (
	"github.com/glamostoffer/ValinorAuth/utils/pg_connector"
	"log/slog"
)

type PgClientRepository struct {
	db    *pg_connector.Connector
	log   *slog.Logger
	Admin AdminRepository
}

func New(db *pg_connector.Connector, log *slog.Logger) *PgClientRepository {
	pg := &PgClientRepository{
		db:  db,
		log: log,
	}

	pg.Admin = newAdminRepo(pg)

	return pg
}
