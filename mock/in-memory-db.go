package mock

import (
	"fmt"
	"github.com/rozturac/cerror"
	"net/http"
	"sync"
)

type InMemoryDB struct {
	collection map[string]string
	locker     *sync.Mutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		collection: make(map[string]string),
		locker:     &sync.Mutex{},
	}
}

func (i *InMemoryDB) Set(key string, value string) error {
	i.locker.Lock()
	defer i.locker.Unlock()
	i.collection[key] = value
	return nil
}

func (i *InMemoryDB) Get(key string) (string, error) {
	i.locker.Lock()
	defer i.locker.Unlock()
	if value, ok := i.collection[key]; ok {
		return value, nil
	} else {
		return "", cerror.NewWithHttpStatusCode(
			cerror.BusinessError,
			fmt.Sprintf("There isn't any value for %s", key),
			http.StatusNotFound,
		)
	}
}
