package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nechel11/Ozon_test/internal/app/store"
	"github.com/sirupsen/logrus"
)

//APIServer структура
type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

func New(config *Config) *APIServer{
	return &APIServer{
		config: config,
		logger : logrus.New(),
		router : mux.NewRouter(),
	}
}


// Start api server
func (s *APIServer) Start () error{
	if err := s.configureLogger(); err != nil{
		return err
	}
	s.configureRouter()

	if err:= s.ConfigureStore(); err != nil{
		return err
	}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}


func (s *APIServer)configureLogger() error{
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil{
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter(){
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) ConfigureStore() error{
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil{
		return err
	}
	s.store = st
	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}