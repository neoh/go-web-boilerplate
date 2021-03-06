package database

import (
    "github.com/boltdb/bolt"
    "github.com/neoh/go-web-boilerplate/models/migrations"
)

var Instance *bolt.DB

func Use(db *bolt.DB) {
	Instance = db
}

func RunMigrations() {
    migrations.BoltDB(Instance)
}