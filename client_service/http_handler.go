package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"client_service/proto"

	"github.com/go-chi/chi/v5"
)

type AuthResponseBody struct {
	Token string `json:"token"`
}

type CreateTaskResponseBody struct {
	TaskID uint64 `json:"task_id"`
}

type Task struct {
	AuthorId         uint64 `json:"author_id"`
	TaskId           uint64 `json:"task_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	DeadlineDate     string `json:"deadline_date"`
	CreationDate     string `json:"creation_date"`
	CompletionStatus string `json:"completion_status"`
}

func SetupHandlers() {
	router := chi.NewRouter()

	router.Post("/register", RegisterHandler)
	router.Post("/auth", AuthHandler)
	router.Put("/update", UpdateHandler)

	router.Post("/task", TaskCreateHandler)
	router.Put("/task", TaskUpdateHandler)
	router.Delete("/task", TaskDeleteHandler)
	router.Get("/task", TaskGetHandler)
	router.Get("/tasks", TasksGetPageHandler)

	router.Post("/task/view", TaskViewHandler)
	router.Post("/task/like", TaskLikeHandler)

	http.Handle("/", router)
}

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("login") || !req.URL.Query().Has("password") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}

	login := req.URL.Query().Get("login")
	tempHash := sha256.Sum256([]byte(req.URL.Query().Get("password")))
	passwordHash := base64.URLEncoding.EncodeToString(tempHash[:])

	if clientDB.LookupLogin(login) != nil {
		http.Error(w, "Login already used", http.StatusForbidden)
		return
	}

	err := clientDB.AddUser(User{Login: login, PasswordHash: passwordHash})
	if err != nil {
		log.Default().Printf("Add user error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	token, err := CreateToken(login)
	if err != nil {
		log.Default().Printf("Create jwt token error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(AuthResponseBody{Token: token})
	if err != nil {
		log.Default().Printf("Response json marshal error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
}

func AuthHandler(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("login") || !req.URL.Query().Has("password") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}

	login := req.URL.Query().Get("login")
	tempHash := sha256.Sum256([]byte(req.URL.Query().Get("password")))
	passwordHash := base64.URLEncoding.EncodeToString(tempHash[:])

	if !clientDB.AuthUser(User{Login: login, PasswordHash: passwordHash}) {
		http.Error(w, "Login or password is incorrent", http.StatusNotFound)
		return
	}

	token, err := CreateToken(login)
	if err != nil {
		log.Default().Printf("Create jwt token error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(AuthResponseBody{Token: token})
	if err != nil {
		log.Default().Printf("Response json marshal error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
}

func UpdateHandler(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("token") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	token := req.URL.Query().Get("token")
	login, err := DecryptToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if clientDB.LookupLogin(login) == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userData, err := clientDB.LoadUserData(login)
	if err != nil {
		log.Default().Printf("Load user data error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&userData)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}

	userData.Login = login

	err = clientDB.UpdateUserData(userData)
	if err != nil {
		log.Default().Printf("Update error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func TaskCreateHandler(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("token") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	token := req.URL.Query().Get("token")
	login, err := DecryptToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := clientDB.LookupLogin(login)
	if userID == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var taskData Task
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&taskData)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	taskData.AuthorId = *userID

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	taskServiceResponse, err := taskService.client.CreateTask(ctx, &proto.CreateTaskRequest{UserId: taskData.AuthorId, Name: taskData.Name, Description: taskData.Description, DeadlineDate: taskData.DeadlineDate})
	if err != nil {
		log.Default().Printf("Create task error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(CreateTaskResponseBody{TaskID: taskServiceResponse.TaskId})
	if err != nil {
		log.Default().Printf("Response json marshal error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
}

func TaskUpdateHandler(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("task_id") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	taskID, err := strconv.ParseUint(req.URL.Query().Get("task_id"), 10, 64)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	if !req.URL.Query().Has("token") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	token := req.URL.Query().Get("token")
	login, err := DecryptToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := clientDB.LookupLogin(login)
	if userID == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var taskData Task
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&taskData)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	taskData.AuthorId = *userID
	taskData.TaskId = taskID

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = taskService.client.UpdateTask(ctx, &proto.UpdateTaskRequest{
		UserId: taskData.AuthorId, TaskId: taskData.TaskId, Name: taskData.Name, Description: taskData.Description, DeadlineDate: taskData.DeadlineDate, CompletionStatus: taskData.CompletionStatus})
	if err != nil {
		log.Default().Printf("Update task error: %v", err)
		http.Error(w, "Task or user not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func TaskDeleteHandler(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("task_id") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	taskID, err := strconv.ParseUint(req.URL.Query().Get("task_id"), 10, 64)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	if !req.URL.Query().Has("token") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	token := req.URL.Query().Get("token")
	login, err := DecryptToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := clientDB.LookupLogin(login)
	if userID == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = taskService.client.DeleteTask(ctx, &proto.DeleteTaskRequest{UserId: *userID, TaskId: taskID})
	if err != nil {
		log.Default().Printf("Delete task error: %v", err)
		http.Error(w, "Task or user not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func TaskGetHandler(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("task_id") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	taskID, err := strconv.ParseUint(req.URL.Query().Get("task_id"), 10, 64)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	if !req.URL.Query().Has("token") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	token := req.URL.Query().Get("token")
	login, err := DecryptToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := clientDB.LookupLogin(login)
	if userID == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	taskServiceResponse, err := taskService.client.GetTask(ctx, &proto.GetTaskRequest{UserId: *userID, TaskId: taskID})
	if err != nil {
		log.Default().Printf("Get task error: %v", err)
		http.Error(w, "Task or user not found", http.StatusNotFound)
		return
	}

	respBody, err := json.Marshal(taskServiceResponse.Task)
	if err != nil {
		log.Default().Printf("Response json marshal error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
}

func TasksGetPageHandler(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("author_login") || !req.URL.Query().Has("page_index") || !req.URL.Query().Has("tasks_per_page") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	authorLogin := req.URL.Query().Get("author_login")
	pageIndex, err := strconv.ParseUint(req.URL.Query().Get("page_index"), 10, 32)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	tasksPerPage, err := strconv.ParseUint(req.URL.Query().Get("tasks_per_page"), 10, 32)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	if !req.URL.Query().Has("token") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	token := req.URL.Query().Get("token")
	login, err := DecryptToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := clientDB.LookupLogin(login)
	if userID == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	authorID := clientDB.LookupLogin(authorLogin)
	if authorID == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	taskServiceResponse, err := taskService.client.GetTasksPage(ctx, &proto.GetTasksPageRequest{UserId: *authorID, PageIndex: uint32(pageIndex), TasksPerPage: uint32(tasksPerPage)})
	if err != nil {
		log.Default().Printf("Get task error: %v", err)
		http.Error(w, "Task or user not found", http.StatusNotFound)
		return
	}

	respBody, err := json.Marshal(taskServiceResponse.Tasks)
	if err != nil {
		log.Default().Printf("Response json marshal error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
}

func TaskViewHandler(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("task_id") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	taskID, err := strconv.ParseUint(req.URL.Query().Get("task_id"), 10, 64)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	if !req.URL.Query().Has("author_id") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	authorID, err := strconv.ParseUint(req.URL.Query().Get("author_id"), 10, 64)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	if !req.URL.Query().Has("token") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	token := req.URL.Query().Get("token")
	login, err := DecryptToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := clientDB.LookupLogin(login)
	if userID == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = statBroker.SendEventMessage(BrokerMessage{UserID: *userID, TaskID: taskID, TaskAuthorId: authorID, EventID: ViewEventID})
	if err != nil {
		log.Printf("Send msg to broker err: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func TaskLikeHandler(w http.ResponseWriter, req *http.Request) {
	if !req.URL.Query().Has("task_id") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	taskID, err := strconv.ParseUint(req.URL.Query().Get("task_id"), 10, 64)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	if !req.URL.Query().Has("author_id") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	authorID, err := strconv.ParseUint(req.URL.Query().Get("author_id"), 10, 64)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	if !req.URL.Query().Has("token") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	token := req.URL.Query().Get("token")
	login, err := DecryptToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := clientDB.LookupLogin(login)
	if userID == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = statBroker.SendEventMessage(BrokerMessage{UserID: *userID, TaskID: taskID, TaskAuthorId: authorID, EventID: LikeEventID})
	if err != nil {
		log.Printf("Send msg to broker err: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
