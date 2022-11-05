package main

import (
	"context"
	"fmt"
	"github.com/simultechnology/my_go_todo_app/config"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	fmt.Println("start web server")

	if err := run(context.Background()); err != nil {
		log.Printf("failed to run server: %v", err)
		os.Exit(1)
	}

	//err := http.ListenAndServe(
	//	":18080",
	//	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	//	}),
	//)
	//if err != nil {
	//	fmt.Printf("failed to terminate server: %v", err)
	//	os.Exit(1)
	//}
	//
	//if err := run(context.Background()); err != nil {
	//	log.Printf("failed to terminate server: %v", err)
	//}

}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port: %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	s := &http.Server{
		//Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			writer := io.MultiWriter(w, os.Stdout)
			fmt.Fprintf(writer, "Hello, %s!!!\n", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// make web server up and running using a goroutine
	eg.Go(func() error {
		if err := s.Serve(l); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdoen: %+v", err)
	}

	return eg.Wait()
}
