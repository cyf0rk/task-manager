package db

import (
	"database/sql"
	"reflect"
	"time"
)

type status int

const (
	TODO status = iota
	INPROGRESS
	DONE
)

func (s status) String() string {
	return [...]string{"todo", "in progress", "done"}[s]
}

type Task struct {
	ID      int
	Name    string
	Project string
	Status  string
	Created time.Time
}

func (t Task) Title() string {
	return t.Name
}

func (t Task) Description() string {
	return t.Project
}

// kancli.Status
func (s status) Next() int {
	if s == DONE {
		return TODO.Int()
	}
	return s.Int() + 1
}

func (s status) Prev() int {
	if s == TODO {
		return DONE.Int()
	}
	return s.Int() - 1
}

func (s status) Int() int {
	return int(s)
}

type TaskDB struct {
	Db *sql.DB
	DataDir string
}

func (t *TaskDB) TableExists(table string) bool {
	stmt, err := t.Db.Prepare("SELECT name FROM sqlite_master WHERE type='table' AND name=?")
	if err != nil {
		return false
	}
	defer stmt.Close()

	var name string
	err = stmt.QueryRow(table).Scan(&name)
	if err != nil {
		return false
	}
	return true
}

func (t *TaskDB) CreateTable() error {
	_, err := t.Db.Exec(`CREATE TABLE "tasks" ("id" INTEGER, "name" TEXT NOT NULL, "project" TEXT,
status TEXT, "created" DATE, PRIMARY KEY("id" AUTOINCREMENT))`)
	return err
}

func (t *TaskDB) Insert(name, project string) error {
	_, err := t.Db.Exec(
		"INSERT INTO tasks (name, project, status, created) VALUES (?, ?, ?, ?)",
		name,
		project,
		TODO.String(),
		time.Now())
	return err
}

func (t *TaskDB) Update(task Task) error {
	// Get the existing state of the task we want to update
	orig, err := t.GetTask(task.ID)
	if err != nil {
		return err
	}
	orig.Merge(task)
	_, err = t.Db.Exec(
		"UPDATE tasks SET name=?, project=?, status=? WHERE id=?",
		orig.Name,
		orig.Project,
		orig.Status,
		orig.ID)
	return err
}

func (orig *Task) Merge(t Task) {
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

func (t *TaskDB) Delete(id int) error {
	_, err := t.Db.Exec("DELETE FROM tasks WHERE id=?", id)
	return err
}

func (t *TaskDB) GetTasks() ([]Task, error) {
	rows, err := t.Db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Name, &t.Project, &t.Status, &t.Created)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (t *TaskDB) GetTasksByStatus(status string) ([]Task, error) {
	rows, err := t.Db.Query("SELECT * FROM tasks WHERE status=?", status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Name, &t.Project, &t.Status, &t.Created)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (t *TaskDB) GetTask(id int) (Task, error) {
	var task Task
	err := t.Db.QueryRow("SELECT * FROM tasks WHERE id=?", id).Scan(&task.ID, &task.Name, &task.Project, &task.Status, &task.Created)
	return task, err
}
