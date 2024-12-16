package main

import (
	"flag"
	"log"
	"web-11/internal/query/api"
	"web-11/internal/query/config"
	"web-11/internal/query/provider"
	"web-11/internal/query/usecase"

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
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessageQuery, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}

// curl -X POST "http://localhost:8083/user?name=User4"
// curl "http://localhost:8083/user?name=Men-fish"