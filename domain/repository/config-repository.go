package repository

import "pure-restfull-api/domain/entity"

type ConfigRepository interface {
	Insert(config *entity.Config) error
	GetByKey(key string) (*entity.Config, error)
}