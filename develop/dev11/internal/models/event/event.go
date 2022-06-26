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

	events := getEvents(store, "-24h")
	return events
}

func (store *EventStore) GetEventsForWeek() []Event {
	store.Lock()
	defer store.Unlock()

	events := getEvents(store, "-168h")
	return events
}

func (store *EventStore) GetEventsForMonth() []Event {
	store.Lock()
	defer store.Unlock()

	events := getEvents(store, "-744h")
	return events
}

func getEvents(store *EventStore, t string) []Event {
	events := make([]Event, 0, len(store.events))
	for _, task := range store.events {
		currTime := time.Now()
		duration, _ := time.ParseDuration(t)
		before := currTime.Add(duration)
		if task.Date.After(before) && task.Date.Before(currTime) {
			events = append(events, task)
		}
	}
	return events
}
