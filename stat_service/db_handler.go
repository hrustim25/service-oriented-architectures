package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	_ "embed"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	DB_URL_ENV = "POSTGRES_URL"
)

type TaskEvent struct {
	TaskId       uint64 `json:"task_id"`
	TaskAuthorId uint64 `json:"task_author_id"`
	UserId       uint64 `json:"user_id"`
	EventID      int    `json:"event_id"`
	EventDate    string `json:"event_date"`
}

type TaskInfo struct {
	TaskId       uint64 `json:"task_id"`
	TaskAuthorId uint64 `json:"task_author_id"`
}

type DBHandler struct {
	db *pgxpool.Pool

	mtx *sync.Mutex
}

var statDB DBHandler

//go:embed sql/create_table.sql
var createTableQuery string

//go:embed sql/add_event.sql
var addEventQuery string

//go:embed sql/get_events.sql
var getEventsQuery string

//go:embed sql/get_event_count.sql
var getEventCountQuery string

//go:embed sql/get_top_tasks.sql
var getTopTasksQuery string

//go:embed sql/get_top_authors.sql
var getTopAuthorsQuery string

func (h *DBHandler) AddEvent(event TaskEvent) error {
	h.mtx.Lock()
	defer h.mtx.Unlock()

	_, err := h.db.Exec(context.Background(), addEventQuery, event.TaskId, event.TaskAuthorId, event.UserId, event.EventID, event.EventDate)
	if err != nil {
		return err
	}
	return nil
}

func (h *DBHandler) GetEvents() ([]TaskEvent, error) {
	h.mtx.Lock()
	defer h.mtx.Unlock()

	var events []TaskEvent
	rows, err := h.db.Query(context.Background(), getEventsQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var event TaskEvent
		err := rows.Scan(&event.TaskId, &event.TaskAuthorId, &event.UserId, &event.EventID, &event.EventDate)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (h *DBHandler) GetEventCount(taskID uint64, eventID uint32) (uint32, error) {
	h.mtx.Lock()
	defer h.mtx.Unlock()

	var count uint32
	err := h.db.QueryRow(context.Background(), getEventCountQuery, taskID, eventID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (h *DBHandler) GetTopTasks(eventID uint32) ([]TaskInfo, error) {
	h.mtx.Lock()
	defer h.mtx.Unlock()

	var tasks []TaskInfo
	rows, err := h.db.Query(context.Background(), getTopTasksQuery, eventID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var taskInfo TaskInfo
		err := rows.Scan(&taskInfo.TaskId, &taskInfo.TaskAuthorId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, taskInfo)
	}
	return tasks, nil
}

func (h *DBHandler) GetTopAuthors() ([]uint64, error) {
	h.mtx.Lock()
	defer h.mtx.Unlock()

	var authorIDs []uint64
	rows, err := h.db.Query(context.Background(), getTopAuthorsQuery, LikeEventID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var authorID uint64
		err := rows.Scan(&authorID)
		if err != nil {
			return nil, err
		}
		authorIDs = append(authorIDs, authorID)
	}
	return authorIDs, nil
}

func SetupDB() {
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

	_, err = db.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table in DB, err: %v", err)
	}

	statDB = DBHandler{db, &sync.Mutex{}}
}
