package api

import (
	"net/http"
)

// NewServer ...
func NewServer() *Server {
	handlers := make(map[Route]func(http.ResponseWriter, *http.Request) (interface{}, *apiErr), NumOfRoutes)
	// add routes
	handlers[Add] = addNumbers
	handlers[Time] = addDaysToDateTime
	server := &Server{
		handle: handlers,
	}
	return server
}
