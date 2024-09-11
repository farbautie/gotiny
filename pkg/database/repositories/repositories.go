package repositories

import "github.com/farbautie/gotiny/pkg/database"

type Repositories struct {
	pool *database.Database
}

func New(pool *database.Database) *Repositories {
	return &Repositories{
		pool: pool,
	}
}
