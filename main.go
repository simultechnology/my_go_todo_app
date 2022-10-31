package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("start web server")

	err := run(context.Background())
	if err != nil {
		fmt.Printf("failed to run server: %v", err)
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

	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// make web server up and running using a goroutine
	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil &&
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
