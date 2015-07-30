package boltdb

import (
	"fmt"
	"path"

	"github.com/boltdb/bolt"
)

const (
	// FIXME: Hardcoded path
	defaultPath = "/tmp/"
)

type DB struct {
	boltDB     *bolt.DB
	dbName     string
	dbPath     string
	bucketName []byte
}

func (db *DB) Open(name string) error {
	var (
		err    error
		boltDB *bolt.DB
	)
	dbPath := path.Join(defaultPath, name)
	boltDB, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}
	var tx *bolt.Tx
	tx, err = boltDB.Begin(true)
	if err != nil {
		return err
	}

	bucketName := []byte(name)
	_, err = tx.CreateBucketIfNotExists(bucketName)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	db.boltDB = boltDB
	db.dbName = name
	db.dbPath = dbPath
	db.bucketName = bucketName
	return nil
}

func (db *DB) Close() error {
	return db.boltDB.Close()
}

func (db *DB) Get(key []byte) ([]byte, error) {
	tx, err := db.boltDB.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	bucket := tx.Bucket(db.bucketName)

	value := bucket.Get(key)

	if value == nil {
		return nil, nil
	}
	return append([]byte{}, value...), nil
}

func (db *DB) Put(key []byte, value []byte) error {
	return db.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(db.bucketName)
		return b.Put(key, value)
	})
}

func (db *DB) String() string {
	return fmt.Sprintf("bolt://%s", db.dbPath)
}

// TODO: map[string]string -> context.Status
func (db *DB) Status() map[string]string {
	// TODO: Do me !
	return map[string]string{
		"driver": "boltdb",
	}
}
