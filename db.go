package main

import (
	"database/sql"
	"reflect"
	"time"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

func (s status) String() string {
	return [...]string{"todo", "in progress", "done"}[s]
}

type task struct {
	ID      int
	Name    string
	Project string
	Status  string
	Created string
}

func (t task) Title() string {
	return t.Name
}

func (t task) Description() string {
	return t.Project
}

// kancli.Status
func (s status) Next() int {
	if s == done {
		return todo.Int()
	}
	return s.Int() + 1
}

func (s status) Prev() int {
	if s == todo {
		return done.Int()
	}
	return s.Int() - 1
}

func (s status) Int() int {
	return int(s)
}

type taskDB struct {
	db *sql.DB
	dataDir string
}

func (t *taskDB) tableExists(table string) bool {
	if _, err := t.db.Exec("SELECT * FROM ?", table); err == nil {
		return true
	}
	return false
}

func (t *taskDB) createTable() error {
	_, err := t.db.Exec(`CREATE TABLE "tasks" ("id" INTEGER, "name" TEXT NOT NULL, "project" TEXT,
status TEXT, "created" DATE, PRIMARY KEY("id" AUTOINCREMENT))`)
	return err
}

func (t *taskDB) insert(name, project string) error {
	_, err := t.db.Exec(
		"INSERT INTO tasks (name, project, status, created) VALUES (?, ?, ?, ?)",
		name,
		project,
		todo.String(),
		time.Now())
	return err
}

func (t *taskDB) update(task task) error {
	// Get the existing state of the task we want to update
	orig, err := t.getTask(task.ID)
	if err != nil {
		return err
	}
	orig.merge(task)
	_, err = t.db.Exec(
		"UPDATE tasks SET name=?, project=?, status=? WHERE id=?",
		orig.Name,
		orig.Project,
		orig.Status,
		orig.ID)
	return err
}

func (orig *task) merge(t task) {
	uValues := reflect.ValueOf(&t).Elem()
	oValues := reflect.ValueOf(orig).Elem()
	for i := 0; i < uValues.NumField(); i++ {
		uField := uValues.Field(i).Interface()
		if oValues.CanSet() {
			if v, ok := uField.(int64); ok && uField != 0 {
				oValues.Field(i).SetInt(v)
			}
			if v, ok := uField.(string); ok && uField != "" {
				oValues.Field(i).SetString(v)
			}
		}
	}
}

func (t *taskDB) delete(id int) error {
	_, err := t.db.Exec("DELETE FROM tasks WHERE id=?", id)
	return err
}

func (t *taskDB) getTasks() ([]task, error) {
	rows, err := t.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task
	for rows.Next() {
		var t task
		err := rows.Scan(&t.ID, &t.Name, &t.Project, &t.Status, &t.Created)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (t *taskDB) getTasksByStatus(status string) ([]task, error) {
	rows, err := t.db.Query("SELECT * FROM tasks WHERE status=?", status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task
	for rows.Next() {
		var t task
		err := rows.Scan(&t.ID, &t.Name, &t.Project, &t.Status, &t.Created)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (t *taskDB) getTask(id int) (task, error) {
	var task task
	err := t.db.QueryRow("SELECT * FROM tasks WHERE id=?", id).Scan(&task.ID, &task.Name, &task.Project, &task.Status, &task.Created)
	return task, err
}
