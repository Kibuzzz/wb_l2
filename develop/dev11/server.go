package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	mux   *http.ServeMux
	store EventStore
}

type errorResponse struct {
	Error string `json:"error"`
}

func New(store EventStore) *Server {
	srv := &Server{}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /create_event", logger(srv.CreateEvent))
	mux.HandleFunc("POST /update_event", srv.UpdateEvent)
	mux.HandleFunc("POST /delete_event", srv.DeleteEvent)

	mux.HandleFunc("GET /events_for_day", srv.EventsForDay)
	mux.HandleFunc("GET /events_for_week", srv.EventsForWeek)
	mux.HandleFunc("GET /events_for_month", srv.EventsForMonth)

	srv.mux = mux
	srv.store = store
	return srv
}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.RequestURI, time.Since(start))
	}
}

func (s *Server) Run() error {
	return http.ListenAndServe(":8080", s.mux)
}

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {

	type createEventRequest struct {
		UserID int   `json:"user_id"`
		Event  Event `json:"event"`
	}

	var req createEventRequest
	err := readBody(r.Body, &req)
	if err != nil {
		renderResponse(errorResponse{Error: err.Error()}, w, http.StatusInternalServerError)
		return
	}

	createdEvent, err := s.store.CreateEvent(req.UserID, req.Event)
	if err != nil {
		renderResponse(errorResponse{err.Error()}, w, http.StatusBadRequest)
		return
	}

	type resp struct {
		Result Event
	}

	renderResponse(resp{Result: createdEvent}, w, http.StatusOK)
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {

	type updateEventRequesr struct {
		UserID int   `json:"user_id"`
		Event  Event `json:"event"`
	}

	var req updateEventRequesr
	err := readBody(r.Body, &req)
	if err != nil {
		renderResponse(errorResponse{err.Error()}, w, http.StatusBadRequest)
		return
	}

	err = s.store.UpdateEvent(req.UserID, req.Event)
	if err != nil {
		renderResponse(errorResponse{err.Error()}, w, http.StatusBadRequest)
		return
	}

	type resp struct {
		Result string
	}

	renderResponse(resp{Result: "event updated successfully"}, w, http.StatusOK)
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {

	type deleteEventRequest struct {
		UserID  int `json:"user_id"`
		EventID int `json:"event_id"`
	}

	var req deleteEventRequest
	err := readBody(r.Body, &req)
	if err != nil {
		renderResponse(errorResponse{err.Error()}, w, http.StatusBadRequest)
		return
	}

	err = s.store.DeleteEvent(req.UserID, req.EventID)
	if err != nil {
		renderResponse(errorResponse{err.Error()}, w, http.StatusBadRequest)
		return
	}

	type resp struct {
		Result string
	}

	renderResponse(resp{Result: "event deleted successfully"}, w, http.StatusOK)
}

func (s *Server) EventsForDay(w http.ResponseWriter, r *http.Request) {

	userID, date, err := parseQueryParams(r)
	if err != nil {
		renderResponse(errorResponse{Error: err.Error()}, w, http.StatusBadRequest)
		return
	}

	events, err := s.store.EventsForDay(userID, date)
	if err != nil {
		renderResponse(errorResponse{err.Error()}, w, http.StatusBadRequest)
		return
	}

	type resp struct {
		Result []Event
	}

	renderResponse(resp{Result: events}, w, http.StatusOK)
}

func (s *Server) EventsForWeek(w http.ResponseWriter, r *http.Request) {

	userID, date, err := parseQueryParams(r)
	if err != nil {
		renderResponse(errorResponse{Error: err.Error()}, w, http.StatusBadRequest)
		return
	}

	events, err := s.store.EventsForWeek(userID, date)
	if err != nil {
		renderResponse(errorResponse{err.Error()}, w, http.StatusBadRequest)
		return
	}

	type resp struct {
		Result []Event
	}

	renderResponse(resp{Result: events}, w, http.StatusOK)
}

func (s *Server) EventsForMonth(w http.ResponseWriter, r *http.Request) {

	userID, date, err := parseQueryParams(r)
	if err != nil {
		renderResponse(errorResponse{Error: err.Error()}, w, http.StatusBadRequest)
		return
	}

	events, err := s.store.EventsForMonth(userID, date)
	if err != nil {
		renderResponse(errorResponse{err.Error()}, w, http.StatusBadRequest)
		return
	}

	type resp struct {
		Result []Event
	}

	renderResponse(resp{Result: events}, w, http.StatusOK)
}

func readBody(body io.ReadCloser, request any) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, &request)
	if err != nil {
		return err
	}
	return nil
}

func renderResponse(resp any, w http.ResponseWriter, code int) {
	w.Header().Add("Content type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func parseQueryParams(r *http.Request) (userID int, date time.Time, err error) {
	userIDstr := r.URL.Query().Get("user_id")
	dateStr := r.URL.Query().Get("date")

	userID, err = strconv.Atoi(userIDstr)
	if err != nil {
		err = fmt.Errorf("invalid user_id: %w", err)
		return
	}

	date, err = time.Parse(time.DateOnly, dateStr)
	if err != nil {
		err = fmt.Errorf("invalid date: %w", err)
		return
	}

	return
}
