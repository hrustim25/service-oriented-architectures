package main

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Task struct {
	AuthorId         uint64 `json:"author_id"`
	TaskId           uint64 `json:"task_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	DeadlineDate     string `json:"deadline_date"`
	CreationDate     string `json:"creation_date"`
	CompletionStatus string `json:"completion_status"`
}

type DBHandler struct {
	db *pgxpool.Pool
}

func LoadQuery(filename string) string {
	result, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("file not found")
	}
	return string(result)
}

func (h *DBHandler) LookupTask(taskId uint64) *uint64 {
	rows, err := h.db.Query(context.Background(), "SELECT author_id FROM tasks WHERE task_id=$1", taskId)
	if err != nil {
		return nil
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	var authorId uint64
	err = rows.Scan(&authorId)
	if err != nil {
		return nil
	}
	return &authorId
}

func (h *DBHandler) CreateTask(task Task) (uint64, error) {
	var taskID uint64
	err := h.db.QueryRow(context.Background(), "INSERT INTO tasks(author_id, name, description, deadline_date, creation_date) VALUES($1, $2, $3, $4, $5) RETURNING task_id",
		task.AuthorId, task.Name, task.Description, task.DeadlineDate, task.CreationDate).Scan(&taskID)
	if err != nil {
		return 0, err
	}
	return taskID, nil
}

func (h *DBHandler) UpdateTask(task Task) error {
	currentTaskAuthorId := h.LookupTask(task.TaskId)
	if currentTaskAuthorId == nil || *currentTaskAuthorId != task.AuthorId {
		return errors.New("unknown task-id and author-id pair")
	}
	_, err := h.db.Exec(context.Background(), "UPDATE tasks SET name=$1, description=$2, deadline_date=$3, completion_status=$4 WHERE task_id=$5",
		task.Name, task.Description, task.DeadlineDate, task.CompletionStatus, task.TaskId)
	return err
}

func (h *DBHandler) DeleteTask(authorID uint64, taskID uint64) error {
	currentTaskAuthorId := h.LookupTask(taskID)
	if currentTaskAuthorId == nil || *currentTaskAuthorId != authorID {
		return errors.New("unknown task-id and author-id pair")
	}
	_, err := h.db.Exec(context.Background(), "DELETE FROM tasks WHERE task_id=$1", taskID)
	return err
}

func (h *DBHandler) GetTask(authorID uint64, taskID uint64) (Task, error) {
	currentTaskAuthorId := h.LookupTask(taskID)
	if currentTaskAuthorId == nil || *currentTaskAuthorId != authorID {
		return Task{}, errors.New("unknown task-id and author-id pair")
	}
	var task Task
	err := h.db.QueryRow(context.Background(), "SELECT task_id, name, description, deadline_date, creation_date, completion_status FROM tasks WHERE task_id=$1",
		taskID).Scan(&task.TaskId, &task.Name, &task.Description, &task.DeadlineDate, &task.CreationDate, &task.CompletionStatus)
	return task, err
}

func (h *DBHandler) GetTasksPage(authorID uint64, pageIndex, pageLimit uint32) ([]Task, error) {
	tasks := make([]Task, 0, pageLimit)
	rows, err := h.db.Query(context.Background(), "SELECT task_id, name, description, deadline_date, creation_date, completion_status FROM tasks WHERE author_id=$1 LIMIT $2 OFFSET $3",
		authorID, pageLimit, (pageIndex-1)*pageLimit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.TaskId, &task.Name, &task.Description, &task.DeadlineDate, &task.CreationDate, &task.CompletionStatus)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func SetupDB() DBHandler {
	if _, ok := os.LookupEnv(DB_URL_ENV); !ok {
		log.Fatalf("'%v' env var not found", DB_URL_ENV)
	}

	time.Sleep(5 * time.Second)

	poolConfig, err := pgxpool.ParseConfig(os.Getenv(DB_URL_ENV))
	if err != nil {
		log.Fatalf("Failed to parse DB config, err: %v", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("Failed to connect to DB, err: %v", err)
	}

	return DBHandler{db}
}
