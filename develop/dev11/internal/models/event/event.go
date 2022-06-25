package event

import (
	"sync"
	"time"
)

type Event struct {
	Id          int       `json:"id"`
	Description string    `json:"text"`
	Date        time.Time `json:"date"`
}

type EventStore struct {
	sync.Mutex
	events map[int]Event
	id     int
}

func NewStore() *EventStore {
	store := &EventStore{}
	store.events = make(map[int]Event)
	store.id = 0
	return store
}

func (store *EventStore) CreateEvent(description string, date time.Time) int {
	store.Lock()
	defer store.Unlock()

	event := Event{
		Id:          store.id,
		Description: description,
		Date:        date,
	}

	store.events[store.id] = event
	store.id++
	return store.id
}

func (store *EventStore) GetEventsForDay() []Event {
	store.Lock()
	defer store.Unlock()

	allTasks := make([]Event, 0, len(store.events))
	for _, task := range store.events {
		allTasks = append(allTasks, task)
	}
	return allTasks
}
