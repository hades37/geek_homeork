package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

// 其中一个ctx取消时能够能够触发另一个退出
func httpServer(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server{Addr: addr, Handler: handler}
	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func serverDebug(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server{Addr: addr, Handler: handler}
	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func main() {
	group, gctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		return httpServer(gctx, "127.0.0.1:8080", http.DefaultServeMux)
	})
	group.Go(func() error {
		return serverDebug(gctx, "127.0.0.1:8090", http.DefaultServeMux)
	})
	// signal 处理,当收到比如kill之类的signal时主动退出
	schan := make(chan os.Signal, 1)
	signal.Notify(schan)
	group.Go(func() error {
		for {
			select {
			case <-gctx.Done():
				return gctx.Err()
			case <-schan:
				gctx.Done()
			}
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		// log
	}
}
