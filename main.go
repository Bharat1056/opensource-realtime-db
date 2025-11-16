package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open(".db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data := map[string]string{
		"Name": "Bharat Panigrahi",
		"Age": "21",
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte("users"))
		if err != nil {
			return err
		}
		for k, v := range data {
			if err := bucket.Put([]byte(k), []byte(v)); err != nil {
				return err
			}
		}
		return nil
	})


	user := make(map[string]string)

	if err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))
		if bucket == nil {
			return fmt.Errorf("bucket (%s) not found", "users")
		}

		bucket.ForEach(func(k, v []byte) error {
			user[string(k)] = string(v)
			return nil
		})

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello world")
}
