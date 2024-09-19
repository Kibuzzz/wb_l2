package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventStore(t *testing.T) {

	date := time.Now()

	t.Run("error when user not found", func(t *testing.T) {
		store := NewStore()
		userID := 0
		dateStart := date
		dateEnd := date
		events, err := store.getEvents(userID, dateStart, dateEnd)
		assert.ErrorIs(t, err, ErrorUserNotFound)
		assert.Empty(t, events)
	})

	t.Run("creating event for user provides no error", func(t *testing.T) {
		store := NewStore()
		userID := 0
		event := Event{Name: "Футбол", Date: date}
		store.CreateEvent(userID, event)
		assert.Equal(t, event, store.events[userID][0])
	})

	t.Run("getting events for existing user within range", func(t *testing.T) {
		store := NewStore()
		userID := 0
		dateStart := date.Add(-time.Hour)
		dateEnd := date.Add(time.Hour)
		event := Event{Date: date}
		store.CreateEvent(userID, event)
		events, err := store.getEvents(userID, dateStart, dateEnd)
		assert.NoError(t, err)
		assert.Len(t, events, 1)
	})

	t.Run("getting events for user outside range returns no events", func(t *testing.T) {
		store := NewStore()
		userID := 0
		dateStart := date.Add(-24 * time.Hour)
		dateEnd := date.Add(-24 * time.Hour)
		event := Event{Date: date}
		store.CreateEvent(userID, event)
		events, err := store.getEvents(userID, dateStart, dateEnd)
		assert.NoError(t, err)
		assert.Empty(t, events)
	})
	t.Run("events for day", func(t *testing.T) {
		store := NewStore()
		userID := 1
		event := Event{Date: date.Add(time.Hour * 6)}
		store.CreateEvent(userID, event)
		events, err := store.EventsForDay(userID, date)
		assert.NoError(t, err)
		assert.NotEmpty(t, events)
	})
	t.Run("events for week", func(t *testing.T) {
		store := NewStore()
		userID := 2
		event := Event{Date: date.Add(time.Hour * 24 * 5)}
		store.CreateEvent(userID, event)
		events, err := store.EventsForWeek(userID, date)
		assert.NoError(t, err)
		assert.NotEmpty(t, events)
	})
	t.Run("events for month", func(t *testing.T) {
		store := NewStore()
		userID := 3
		event := Event{Date: date.Add(time.Hour * 24 * 10)}
		store.CreateEvent(userID, event)
		events, err := store.EventsForMonth(userID, date)
		assert.NoError(t, err)
		assert.NotEmpty(t, events)
	})

	t.Run("update event", func(t *testing.T) {
		store := NewStore()
		userID := 4
		event := Event{Name: "Праздник", Date: date}
		createdEvent, err := store.CreateEvent(userID, event)
		assert.NoError(t, err)

		updatedEvent := createdEvent
		updatedEvent.Name = "Корпоратив"
		updatedEvent.Date = updatedEvent.Date.Add(time.Hour)
		err = store.UpdateEvent(userID, updatedEvent)
		assert.NoError(t, err)
	})
	t.Run("update wrong event", func(t *testing.T) {
		store := NewStore()
		userID := 4
		event := Event{Name: "Праздник", Date: date}
		err := store.UpdateEvent(userID, event)
		assert.ErrorIs(t, err, ErrorUserNotFound)
		event, err = store.CreateEvent(userID, event)
		assert.NoError(t, err)
		event.ID = -1 // неправильны ID
		err = store.UpdateEvent(userID, event)
		assert.ErrorIs(t, err, ErrorUserNotFound)
	})
}
