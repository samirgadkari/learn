package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/samirgadkari/learn/web/basic/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sig_chan := make(chan os.Signal)
	signal.Notify(sig_chan, os.Interrupt)
	signal.Notify(sig_chan, os.Kill)

	sig := <-sig_chan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
