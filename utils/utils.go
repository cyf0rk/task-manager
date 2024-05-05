package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"task-manager/db"
	_ "github.com/mattn/go-sqlite3"
	gap "github.com/muesli/go-app-paths"
)

func SetupPath() string {
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

func OpenDB(path string) (*db.TaskDB, error) {
	sqlDb, err := sql.Open("sqlite3", filepath.Join(path, "tasks.db"))
	if err != nil {
		return nil, err
	}
	t := db.TaskDB{Db: sqlDb, DataDir: path}
	if !t.TableExists("tasks") {
		fmt.Println("Creating table")
		err := t.CreateTable()
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}
