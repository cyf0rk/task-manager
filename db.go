package main

import (
	"database/sql"
)

type taskDB struct {
	db *sql.DB
	dataDir string
}

func (t *taskDB) tableExists(table string) bool {
	if _, err := t.db.Query("SELECT * FROM tasks"); err == nil {
		return true
	}
	return false
}

func (t *taskDB) createTable() error {
	_, err := t.db.Exec(`CREATE TABLE "tasks" ("id" INTEGER, "name" TEXT NOT NULL, "project" TEXT,
status TEXT, "created" DATE, PRIMARY KEY("id" AUTOINCREMENT))`)
	return err
}
