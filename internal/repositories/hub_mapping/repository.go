package hub_mapping_repo

import (
	"sync"

	"github.com/omniful/go_commons/db/sql/postgres"
)

// Repository handles queries for order partners.
type Repository struct {
	db *postgres.DbCluster
}

var (
	repository     *Repository
	repositoryOnce sync.Once
)

// NewRepository initializes and returns a singleton instance of Repository.
func NewRepository(db *postgres.DbCluster) *Repository {
	repositoryOnce.Do(func() {
		repository = &Repository{
			db: db,
		}
	})
	return repository
}
