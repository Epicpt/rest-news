package httpserver

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App             *fiber.App
	notify          chan error
	address         string
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration
	prefork         bool
}

func New(port string) *Server {
	s := &Server{
		notify:          make(chan error, 1),
		address:         ":" + port,
		readTimeout:     5 * time.Second,
		writeTimeout:    5 * time.Second,
		shutdownTimeout: 10 * time.Second,
		prefork:         false,
	}

	app := fiber.New(fiber.Config{
		Prefork:      s.prefork,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	s.App = app

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.App.Listen(s.address)
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	return s.App.ShutdownWithTimeout(s.shutdownTimeout)
}
