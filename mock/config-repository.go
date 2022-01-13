package mock

import (
	"pure-restfull-api/domain/entity"
)

type configRepository struct {
	db *InMemoryDB
}

func NewConfigRepository(db *InMemoryDB) *configRepository {
	return &configRepository{
		db: db,
	}
}

func (c configRepository) Insert(config *entity.Config) error {
	return c.db.Set(config.GetKey(), config.GetValue())
}

func (c configRepository) GetByKey(key string) (*entity.Config, error) {
	if value, err := c.db.Get(key); err != nil {
		return nil, err
	} else {
		return entity.CreateConfig(key, value), nil
	}
}
