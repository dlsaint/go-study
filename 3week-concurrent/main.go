package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	//利用errgroup特性, 确保服务同时存在或退出
	g, ctx := errgroup.WithContext(context.Background())
	wg := sync.WaitGroup{}
	server := NewHttpServer() //http服务
	//退出监听
	g.Go(func() error {
		<-ctx.Done()
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		log.Println("shutting down server...")
		return server.Shutdown(timeoutCtx)
	})

	//启动http服务
	wg.Add(1)
	g.Go(func() error {
		wg.Done()
		return server.ListenAndServe()
	})
	wg.Wait()

	//监听信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-c:
				return errors.Errorf("get os signal: %v", sig)
			}
		}
	})

	fmt.Printf("exiting, err is: %+v\n", g.Wait())
}

//新建http服务
func NewHttpServer() http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	return http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
}
