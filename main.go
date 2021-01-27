package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/shyshlakov/go-http-server/handler"
	"github.com/shyshlakov/go-http-server/persistence/repo"
	"github.com/shyshlakov/go-http-server/server"
	"github.com/shyshlakov/go-http-server/service"
)

func main() {
	repo := &repo.PostgreRepo{}
	service := &service.AppService{
		Repo: repo,
	}
	handler := handler.NewServerRoutes(service)
	serv := server.NewAppServer(handler)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT)
	select {
	case err := <-serv.ListenAndServe():
		fmt.Println(err)
	case <-sigCh:
		fmt.Println("ok")
	}
}
