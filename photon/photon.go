package photon

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

const (
	defaultDbName = "default"
)

type Collection struct {}

type Photon struct {
	db *bolt.DB
}

func New() (*Photon, error) {
	dbName := fmt.Sprintf("%s.photon", defaultDbName)
	db, err := bolt.Open(dbName, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &Photon{
		db: db,
	}, nil
}

func (p *Photon) CreateCollection(name string) (*Collection, error) {

}
