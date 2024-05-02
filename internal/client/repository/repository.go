package repository

import (
	"github.com/glamostoffer/ValinorAuth/utils/pg_connector"
	"log/slog"
)

type PgClientRepository struct {
	db   *pg_connector.Connector
	log  *slog.Logger
	User UserRepository
}

func New(db *pg_connector.Connector, log *slog.Logger) *PgClientRepository {
	pg := &PgClientRepository{
		db:  db,
		log: log,
	}

	pg.User = newUserRepo(pg)

	return pg
}
