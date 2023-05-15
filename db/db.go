package db

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/spf13/viper"
)

type KVDatabase struct {
	kv *badger.DB
}

func New() (*KVDatabase, error) {
	db, err := badger.Open(badger.DefaultOptions(viper.GetString("db.path")))
	if err != nil {
		return nil, err
	}
	return &KVDatabase{kv: db}, nil
}

func (db *KVDatabase) Close() {
	db.kv.Close()
}
