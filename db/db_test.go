package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetTask(t *testing.T) {
	tests := []struct {
		want Task
	}{
		{
			want: Task{
				ID:      1,
				Name:    "finish script",
				Project: "development",
				Status:  todo.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run("GetTask", func(t *testing.T) {
			db := setup()
			defer teardown(db)
			if err := db.Insert(tt.want.Name, tt.want.Project); err != nil {
				t.Fatalf("We run into unexpected error: %v", err)
			}
			task, err := db.GetTask(tt.want.ID)
			if err != nil {
				t.Fatalf("We run into unexpected error: %v", err)
			}
			tt.want.Created = task.Created
			if !reflect.DeepEqual(task, tt.want) {
				t.Errorf("getTask() got = %v, want %v", task, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		want Task
	}{
		{
			want: Task{
				ID:      1,
				Name:    "finish script",
				Project: "development",
				Status:  todo.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run("GetTask", func(t *testing.T) {
			db := setup()
			defer teardown(db)
			if err := db.Insert(tt.want.Name, tt.want.Project); err != nil {
				t.Fatalf("We run into unexpected error: %v", err)
			}
			tasks, err := db.GetTasks()
			if err != nil {
				t.Fatalf("We run into unexpected error: %v", err)
			}
			tt.want.Created = tasks[0].Created
			if !reflect.DeepEqual(tasks[0], tt.want) {
				t.Errorf("getTask() got = %v, want %v", tasks, tt.want)
			}
			if err := db.Delete(tt.want.ID); err != nil {
				t.Fatalf("Unable to delete tasks: %v", err)
			}
			tasks, err = db.GetTasks()
			if err != nil {
				t.Fatalf("Unable to get tasks: %v", err)
			}
			if len(tasks) != 0 {
				t.Fatalf("Expected tasks to be empty, got %d", len(tasks))
			}
		})
	}
}

func TestGetTasksByStatus(t *testing.T) {
	tests := []struct {
		want Task
	}{
		{
			want: Task{
				ID:      1,
				Name:    "finish script",
				Project: "development",
				Status:  todo.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run("GetTaskByStatus", func(t *testing.T) {
			db := setup()
			defer teardown(db)
			if err := db.Insert(tt.want.Name, tt.want.Project); err != nil {
				t.Fatalf("We run into unexpected error: %v", err)
			}
			tasks, err := db.GetTasksByStatus(tt.want.Status)
			if err != nil {
				t.Fatalf("We run into unexpected error: %v", err)
			}
			if len(tasks) < 1 {
				t.Fatalf("Expected 1 task, got %d", len(tasks))
			}
			tt.want.Created = tasks[0].Created
			if !reflect.DeepEqual(tasks[0], tt.want) {
				t.Errorf("getTasksByStatus() got = %v, want %v", tasks, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		want Task
	}{
		{
			want: Task{
				ID:      1,
				Name:    "finish script",
				Project: "development",
				Status:  todo.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run("Update", func(t *testing.T) {
			db := setup()
			defer teardown(db)
			if err := db.Insert(tt.want.Name, tt.want.Project); err != nil {
				t.Fatalf("We run into unexpected error: %v", err)
			}
			tasks, err := db.GetTasks()
			if err != nil {
				t.Fatalf("We run into unexpected error: %v", err)
			}
			tt.want.Created = tasks[0].Created
			if !reflect.DeepEqual(tasks[0], tt.want) {
				t.Errorf("getTasks() got = %v, want %v", tasks, tt.want)
			}
			tasks[0].Status = done.String()
			if err := db.Update(tasks[0]); err != nil {
				t.Fatalf("Unable to update tasks: %v", err)
			}
			tasks, err = db.GetTasks()
			if err != nil {
				t.Fatalf("Unable to get tasks: %v", err)
			}
			if tasks[0].Status != done.String() {
				t.Fatalf("Expected task status to be done, got %s", tasks[0].Status)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		want Task
	}{
		{
			name: "merge name",
			want: Task{
				ID:      1,
				Name:    "finish script",
				Project: "development",
				Status:  todo.String(),
			},
		},
		{
			name: "merge project",
			want: Task{
				ID:      1,
				Name:    "finish script",
				Project: "development",
				Status:  todo.String(),
			},
		},
		{
			name: "merge status",
			want: Task{
				ID:      1,
				Name:    "finish script",
				Project: "development",
				Status:  todo.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var orig Task
			orig.ID = 1
			orig.Name = "write script"
			orig.Project = "development"
			orig.Status = todo.String()
			orig.Created = "2020-01-01 00:00:00"
			tt.want.Created = "2020-01-01 00:00:00"
			orig.Merge(tt.want)
			if !reflect.DeepEqual(orig, tt.want) {
				t.Errorf("merge() got = %v, want %v", orig, tt.want)
			}
		})
	}
}

func setup() *TaskDB {
	path := filepath.Join(os.TempDir(), "test.db")
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	t := TaskDB{db, path}
	if !t.TableExists("tasks") {
		if err := t.CreateTable(); err != nil {
			log.Fatal(err)
		}
	}
	return &t
}

func teardown(t *TaskDB) {
	t.Db.Close()
	os.Remove(t.DataDir)
}
