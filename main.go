package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	gap "github.com/muesli/go-app-paths"
)

func setupPath() string {
	scope := gap.NewScope(gap.User, "tasks")
	dirs, err := scope.DataDirs()
	if err != nil {
		log.Fatal(err)
	}
	// Create the directory if it doesn't exist
	var taskDir string
	if len(dirs) > 0 {
		taskDir = dirs[0]
	} else {
		taskDir, _ = os.UserHomeDir()
	}
	if err := initTaskDir(taskDir); err != nil {
		log.Fatal(err)
	}
	return taskDir
}

func initTaskDir(taskDir string) error {
	if _, err := os.Stat(taskDir); err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(taskDir, 0o750)
		}
		return err
	}
	return nil
}

func openDB(path string) (*taskDB, error) {
	db, err := sql.Open("sqlite3", filepath.Join(path, "tasks.db"))
	if err != nil {
		return nil, err
	}
	t := taskDB{db: db, dataDir: path}
	if !t.tableExists("tasks") {
		fmt.Println("Creating table")
		err := t.createTable()
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}

func main() {
	return
}
