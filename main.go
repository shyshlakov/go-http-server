package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/shyshlakov/go-http-server/config"
	"github.com/shyshlakov/go-http-server/handler"
	"github.com/shyshlakov/go-http-server/persistence/repo/postgres"
	"github.com/shyshlakov/go-http-server/restapi"
	"github.com/shyshlakov/go-http-server/server"
	"github.com/shyshlakov/go-http-server/service"
)

func main() {
	swagger, err := restapi.GetSwagger()
	if err != nil {
		log.Fatal("Can not get swagger spec for api")
	}
	cfg, err := config.FromEnv()
	if err != nil {
		fmt.Printf("Can not get env: %v", err)
		return
	}
	repo := postgres.NewRepo(cfg)
	if err := repo.Connect(); err != nil {
		fmt.Printf("Can not connect to DB: %v", err)
		return
	}
	defer func() {
		if err := repo.Close(); err != nil {
			fmt.Printf("Can not close connect to DB: %v", err)
		}
	}()
	service := service.NewService(repo)
	handler := handler.NewServerRoutes(service)
	serv := server.NewAppServer(handler, cfg, swagger)
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
