package db

import "sync"

type DB struct {
	data map[string]string
	mu   sync.Mutex
}

func New() *DB {
	return &DB{
		data: make(map[string]string),
	}
}

func (db *DB) Set(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

func (db *DB) Get(key string) string {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.data[key]
}

func (db *DB) Del(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, key)
}
