package photon

import (
	"fmt"

	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

const (
	defaultDbName = "default"
)

type M map[string]string

type Collection struct {
	*bolt.Bucket
}

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
	tx, err := p.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		return  nil, err
	}

	return &Collection{
		Bucket: bucket,
	}, nil
}

func (p *Photon) Insert(collName string, data M) (uuid.UUID, error) {
		id := uuid.New()

		tx, err := p.db.Begin(true)
		if err != nil {
			return id, err
		}
		defer tx.Rollback()

		bucket, err := tx.CreateBucketIfNotExists([]byte(collName))
		if err != nil {
			return  id, err
		}

		for k, v := range data {
			if err := bucket.Put([]byte(k), []byte(v)); err != nil {
				return id, err
			}
		}

		if err := bucket.Put([]byte("id"), []byte(id.String())); err != nil {
			return id, err
		}

		return id, tx.Commit()
}

func (p *Photon) Select(coll, k string, query M) (M, error){
	tx, err := p.db.Begin(false)
	if err != nil {
		return nil, err
	}
	bucket := tx.Bucket([]byte(coll))
	if bucket == nil {
		return nil, fmt.Errorf("Collection (%s) not found", coll)
	}
}
