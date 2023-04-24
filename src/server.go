package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ellofae/DNS-searcher-Web-Server/src/dataProccess"

	"github.com/ellofae/DNS-searcher-Web-Server/src/handlers"
)

// Type data definitions

func main() {
	err := dataProccess.LoadAllData()
	if err != nil {
		fmt.Println(err)
		return
	}

	l := log.New(os.Stdout, "DNS-service", log.LstdFlags)
	homePageHandler := handlers.NewHomePage(l)
	getIPHandler := handlers.NewGetIP(l)
	getNameHandler := handlers.NewGetName(l)
	getNameServersHandler := handlers.NewGetNameServers(l)
	getMailServersHandler := handlers.NewGetMailServers(l)

	sm := http.NewServeMux()
	sm.Handle("/", homePageHandler)
	sm.Handle("/getIP", getIPHandler)
	sm.Handle("/getName", getNameHandler)
	sm.Handle("/nameServers", getNameServersHandler)
	sm.Handle("/mailServers", getMailServersHandler)

	srv := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recived terminate, gracefil shutdown", sig)

	// Graceful shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)
}
