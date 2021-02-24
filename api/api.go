package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Route ...
type Route string

const (
	// Hello ...
	Hello Route = "/api/hello"
	// Add ...
	Add Route = "/api/add"
	// Hello ...
	Time Route = "/api/time"
	// NumOfRoutes ...
	NumOfRoutes = 2
)

type (
	// Server ...
	Server struct {
		handle map[Route]func(http.ResponseWriter, *http.Request) (interface{}, *apiErr)
	}
)

// ServeHTTP...
func (s *Server) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	switch Route(req.RequestURI) {
	case Hello:
		if req.Method == "POST" {
			json.NewEncoder(writer).Encode(newNotImplemented(req.Method, req.RequestURI))
			return
		}
		fmt.Fprintf(writer, "hello\n")

	case Add:
		if req.Method == "GET" {
			json.NewEncoder(writer).Encode(newNotImplemented(req.Method, req.RequestURI))
			return
		}
		apiResp, err := s.handle[Add](writer, req)
		if err != nil {
			json.NewEncoder(writer).Encode(err)
			return
		}
		sum := apiResp.(Sum)

		fmt.Fprintf(writer, fmt.Sprintf("%d\n", sum.Total))

	case Time:
		if req.Method == "GET" {
			json.NewEncoder(writer).Encode(newNotImplemented(req.Method, req.RequestURI))
			return
		}
		apiResp, err := s.handle[Time](writer, req)
		if err != nil {
			json.NewEncoder(writer).Encode(err)
			return
		}
		days := apiResp.(DateTime)

		fmt.Fprintf(writer, fmt.Sprintf("%s\n", days.IncrementedDateTime))
	}
}

func addNumbers(w http.ResponseWriter, r *http.Request) (interface{}, *apiErr) {
	var numbers Numbers
	data, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(data, &numbers)
	if err != nil {
		return nil, newInternalServerErr(err)
	}
	// Input Validation
	if validationErrs := numbers.validate(); len(validationErrs) > 0 {
		return nil, newInvalidArgumentErr(validationErrs)
	}
	return Sum{*numbers.First + *numbers.Second}, nil
}

func addDaysToDateTime(w http.ResponseWriter, r *http.Request) (interface{}, *apiErr) {
	var days Days
	data, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(data, &days)
	if err != nil {
		return nil, newInternalServerErr(err)
	}
	// Input Validation
	if validationErrs := days.validate(); len(validationErrs) > 0 {
		return nil, newInvalidArgumentErr(validationErrs)
	}
	return DateTime{time.Now().AddDate(0, 0, *days.NumberOfDays)}, nil
}
