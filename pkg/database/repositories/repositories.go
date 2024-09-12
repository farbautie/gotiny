package repositories

import (
	"database/sql"

	"github.com/farbautie/gotiny/pkg/database"
)

type Repositories struct {
	pool *sql.DB
}

func New(pool *database.Database) *Repositories {
	return &Repositories{
		pool: pool.Pool,
	}
}
