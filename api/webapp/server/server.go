package server

import (
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp"
	"microseviceAdmin/webapp/logger"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	config *webapp.Config
	logger *logger.Logger
	router *httprouter.Router
	Store  *store.Store
}

// New ...
func New(config *webapp.Config) *Server {
	return &Server{
		config: config,
		logger: logger.GetLogger(),
		router: httprouter.New(),
	}
}

// Start ...
func (s *Server) Start() error {

	s.configureRoutes()
	s.logger.Info("Admin router started successfully")

	if err := s.configureStore(); err != nil {
		s.logger.Errorf("Error while configure store. ERR MSG: %s", err.Error())
		return err
	}
	s.logger.Info("Store started successfully")

	s.logger.Infof("Server starts at %s ...", s.config.ServerInfo())

	return http.ListenAndServe(s.config.ServerAddress(), s.router)
}
