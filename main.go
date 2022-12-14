package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/simultechnology/my_go_todo_app/config"
	"github.com/simultechnology/my_go_todo_app/server"
)

func main() {
	fmt.Println("start web server")

	if err := run(context.Background()); err != nil {
		log.Printf("failed to run server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		// log.Fatalf("failed to listen port: %d: %v", cfg.Port, err)
		log.Panicf("failed to listen port: %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	mux := server.NewMux()
	s := server.NewServer(l, mux)
	return s.Run(ctx)
}
