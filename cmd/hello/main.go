package main

import (
	"flag"
	"log"
	"web-11/internal/hello/api"
	"web-11/internal/hello/config"
	"web-11/internal/hello/provider"
	"web-11/internal/hello/usecase"

	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "../../configs/config.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessage, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}

// curl -X POST -H "Content-Type: application/json" -d '{"msg": "Привет, мир!"}' http://127.0.0.1:8081/hello
// curl -X POST -H "Content-Type: application/json" -d '{"msg": "мир!"}' http://127.0.0.1:8081/hello
// curl -X GET http://127.0.0.1:8081/hello