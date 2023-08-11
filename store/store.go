package store

type DB struct {
	Name string
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Test1() string {
	return "store worked"
}
