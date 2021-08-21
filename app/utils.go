package app

import (
	"fmt"

	dbm "github.com/tendermint/tm-db"
)

var (
	// This is set at compile time. Could be cleveldb, defaults is goleveldb.
	DBBackend = ""
	backend   = dbm.MongoDBBackend
)

func init() {
	if len(DBBackend) != 0 {
		backend = dbm.BackendType(DBBackend)
	}
}

// NewLevelDB instantiate a new LevelDB instance according to DBBackend.
func NewMongoDB(name, dir string) (db dbm.DB, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("couldn't create db: %v", r)
		}
	}()

	return dbm.NewDB(name, backend, dir)
}
