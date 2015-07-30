package storage

import (
	"fmt"

	"github.com/mouadino/metastore/storage/boltdb"
)

func Init(driverName, DBName string) (DB, error) {
	db, err := getDriverDB(driverName)
	if err != nil {
		return nil, err
	}
	err = db.Open(DBName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDriverDB(name string) (DB, error) {
	switch {
	case name == "boltdb":
		return &boltdb.DB{}, nil
	default:
		return nil, fmt.Errorf("unknown storage driver %s", name)
	}
}
