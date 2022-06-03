package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer
type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Route
}

// NEW ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//START ...
func(s *APIServer) Start() error{
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BinAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandlerFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	type request struct{
		name string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Heloo")
	}
}