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

		errUpdate := b.Put([]byte("2015-01-02"), []byte("My New Year post Young 2"))
		b.Put([]byte("2015-01-03"), []byte("My New Year post Young 3"))
		b.Put([]byte("2015-01-04"), []byte("My New Year post Young 4"))
		return errUpdate
	})

	/*	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		v := b.Get([]byte("2015-01-02"))
		fmt.Printf("%sn", v)
		return nil
	})*/

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("posts"))

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})

	defer db.Close()

}
