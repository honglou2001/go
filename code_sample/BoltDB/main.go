package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("blog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("posts"))
		if err != nil {
			return err
		}
		return b.Put([]byte("2015-01-02"), []byte("My New Year post Young"))
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		v := b.Get([]byte("2015-01-02"))
		fmt.Printf("%sn", v)
		return nil
	})


	defer db.Close()

}