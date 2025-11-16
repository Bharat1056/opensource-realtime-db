package photon

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

const (
	defaultDbName = "default"
)

type Collection struct {
	bucket *bolt.Bucket
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
	coll := Collection{}
	err := p.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte("users"))
		if err != nil {
			return  err
		}
		coll.bucket = bucket
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &coll, nil
}


	// user := make(map[string]string)

	// if err := db.View(func(tx *bolt.Tx) error {
	// 	bucket := tx.Bucket([]byte("users"))
	// 	if bucket == nil {
	// 		return fmt.Errorf("bucket (%s) not found", "users")
	// 	}

	// 	bucket.ForEach(func(k, v []byte) error {
	// 		user[string(k)] = string(v)
	// 		return nil
	// 	})

	// 	return nil
	// }); err != nil {
	// 	log.Fatal(err)
	// }
