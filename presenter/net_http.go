package presenter

import (
	"context"
	"github.com/SananGuliyev/gossignment/application/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type NetHttpServer struct {
	inMemoryHandler *handler.InMemoryHandler
	recordHandler   *handler.RecordHandler
}

func NewNetHttpServer(
	inMemoryHandler *handler.InMemoryHandler,
	recordHandler *handler.RecordHandler,
) Server {
	return &NetHttpServer{
		inMemoryHandler: inMemoryHandler,
		recordHandler:   recordHandler,
	}
}

func (s *NetHttpServer) Run() {
	wait := time.Second * 15

	srv := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: s,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()
	<-ctx.Done()

	log.Println("Shutting down")
	os.Exit(0)
}

func (s *NetHttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	switch {
	case p == "/records/filter" && r.Method == http.MethodPost:
		s.recordHandler.Filter(w, r)
	case p == "/in-memory" && r.Method == http.MethodPost:
		s.inMemoryHandler.Create(w, r)
	case p == "/in-memory" && r.Method == http.MethodGet:
		s.inMemoryHandler.Read(w, r)
	default:
		http.NotFound(w, r)
	}
}
