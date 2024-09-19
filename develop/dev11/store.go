package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrorUserNotFound  = errors.New("user not found")
	ErrorEventNotFound = errors.New("event not found")
)

var (
	week  = time.Hour * 24 * 7
	month = time.Hour * 24 * 30
)

type Event struct {
	ID   int
	Name string
	Date time.Time
}

type EventStore interface {
	CreateEvent(userID int, event Event) (Event, error)
	UpdateEvent(userID int, event Event) error
	DeleteEvent(userID int, eventID int) error

	EventsForDay(userID int, date time.Time) ([]Event, error)
	EventsForWeek(userID int, date time.Time) ([]Event, error)
	EventsForMonth(userID int, date time.Time) ([]Event, error)
}

type Store struct {
	events map[int]map[int]Event
	id     int
}

func NewStore() *Store {
	return &Store{events: make(map[int]map[int]Event)}
}

func (es *Store) CreateEvent(userID int, event Event) (Event, error) {
	event.ID = es.id
	if _, ok := es.events[userID]; !ok {
		es.events[userID] = make(map[int]Event)
	}
	es.events[userID][event.ID] = event
	es.id++
	return event, nil
}

func (es *Store) UpdateEvent(userID int, updateEvent Event) error {
	events, ok := es.events[userID]
	if !ok {
		return ErrorUserNotFound
	}
	if _, ok := events[updateEvent.ID]; !ok {
		return ErrorEventNotFound
	}
	events[updateEvent.ID] = updateEvent
	return nil
}

func (es *Store) DeleteEvent(userID int, eventID int) error {
	events, ok := es.events[userID]
	if !ok {
		return ErrorUserNotFound
	}
	if _, ok := events[eventID]; !ok {
		return ErrorEventNotFound
	}
	delete(events, eventID)
	return nil
}

func (es *Store) EventsForDay(userID int, date time.Time) ([]Event, error) {
	return es.getEvents(userID, date, date)
}

func (es *Store) EventsForWeek(userID int, date time.Time) ([]Event, error) {
	return es.getEvents(userID, date, date.Add(week))
}

func (es *Store) EventsForMonth(userID int, date time.Time) ([]Event, error) {
	return es.getEvents(userID, date, date.Add(month))
}

func (es *Store) getEvents(userID int, dateStart time.Time, dateEnd time.Time) ([]Event, error) {
	dateStart = normalizeDate(dateStart)
	dateEnd = normalizeDate(dateEnd)

	events, ok := es.events[userID]
	if !ok {
		return nil, ErrorUserNotFound
	}
	var res []Event
	for _, event := range events {
		eventDate := normalizeDate(event.Date)
		fmt.Println(dateStart, eventDate, dateEnd, (dateStart.Before(eventDate) || dateStart.Equal(eventDate)) && (dateEnd.After(eventDate) || dateEnd.Equal(eventDate)))
		if (dateStart.Before(eventDate) || dateStart.Equal(eventDate)) && (dateEnd.After(eventDate) || dateEnd.Equal(eventDate)) {
			res = append(res, event)
		}
	}
	return res, nil
}

func normalizeDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}
