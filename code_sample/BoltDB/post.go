package main

import (
	"time"
)

type Post struct {
	Created time.Time
	Title   string
	Content string
}


/*post := &Post{
Created: time.Now(),
Title:   "My first post",
Content: "Hello, this is my first post.",
}

db.Update(func(tx *bolt.Tx) error {
	b, err := tx.CreateBucketIfNotExists([]byte("posts"))
	if err != nil {
		return err
	}
	encoded, err := json.Marshal(post)
	if err != nil {
		return err
	}
	return b.Put([]byte(post.Created.Format(time.RFC3339)), encoded)
})*/