package main

import (
	"context"
	"fmt"
	"github.com/witchcraft1/Architecture-course-2/server/dormitories"
	"net/http"
)

type HttpPortNumber int

type DormitoryApiServer struct {
	Port HttpPortNumber
	DormitoriesHandler dormitories.HttpHandlerFunc
	server *http.Server
}

func (s *DormitoryApiServer) Start() error {
	if s.DormitoriesHandler == nil {
		return fmt.Errorf("Dormitories HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/postData", s.DormitoriesHandler)
	handler.HandleFunc("/getData", s.DormitoriesHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *DormitoryApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
