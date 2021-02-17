package addressstorage

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type AddressDB struct {
	db *leveldb.DB
}

func (l *AddressDB) open(dbpath string) error {

	var err error
	l.db, err = leveldb.OpenFile(dbpath, nil)
	return err

}

func NewDB(dbpath string) (*AddressDB, error) {
	db := &AddressDB{}
	err := db.open(dbpath)

	return db, err
}

func (l *AddressDB) Close() {
	l.db.Close()

}
func (l *AddressDB) HasValue(key string) bool {

	_, err := l.db.Get([]byte(key), nil)

	if err != nil {
		return false
	}
	return true
}

func (l *AddressDB) SetValue(key string) error {
	return l.db.Put([]byte(key), []byte(key), nil)
}

func (l *AddressDB) BatchWrite(kv []string) error {

	batch := new(leveldb.Batch)
	for i := 0; i < len(kv); {
		batch.Put([]byte(kv[i]), []byte(kv[i+1]))
		i += 1
	}

	if err := l.db.Write(batch, nil); err != nil {
		return err
	}

	return nil
}
