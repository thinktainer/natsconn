package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"net/http"
	_ "net/http/pprof"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats-streaming"
)

func main() {
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)
	go servePProf()
	for i := 0; i < 100; i++ {
		log.Println("connect", i)
		nc, err := nats.Connect("nats://localhost:4222")
		if err != nil {
			log.Fatal(err)
		}
		conn, err := stan.Connect("test-cluster", "nats-test", stan.NatsConn(nc))
		if err != nil {
			log.Fatal(err)
		}
		sub, err := conn.Subscribe("test-subject", func(*stan.Msg) {
			return
		})
		if err != nil {
			log.Fatal(err)
		}
		err = sub.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		if conn.NatsConn() != nil && !conn.NatsConn().IsClosed() {
			conn.NatsConn().Close()
		} else {
			log.Println("nats conn was no longer present")
		}
	}
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)
	runtime.Goexit()
}

func servePProf() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}
	_ = srv.ListenAndServe()
}
