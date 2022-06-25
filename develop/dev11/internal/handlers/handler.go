package handlers

import (
	"encoding/json"
	"eventAPI/internal/models/event"
	"fmt"
	"log"
	"mime"
	"net/http"
	"time"
)

type eventServer struct {
	store *event.EventStore
}

func NewTaskServer() *eventServer {
	store := event.NewStore()
	return &eventServer{store: store}
}

func (s *eventServer) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		type RequestEvent struct {
			Id          int       `json:"id"`
			Description string    `json:"description"`
			Date        time.Time `json:"date"`
		}
		type ResponseId struct {
			Id int `json:"id"`
		}
		contentType := r.Header.Get("Content-Type")
		mediatype, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if mediatype != "application/json" {
			http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
			return
		}

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var rt RequestEvent
		if err := dec.Decode(&rt); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := s.store.CreateEvent(rt.Description, rt.Date)
		js, err := json.Marshal(ResponseId{Id: id})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func (s *eventServer) GetEventsForDay(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get all tasks at %s\n", req.URL.Path)

	allTasks := s.store.GetEventsForDay()

	js, err := json.Marshal(allTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(js)
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(js)
	if err != nil {
		log.Fatalln(write, err)
	}
}
