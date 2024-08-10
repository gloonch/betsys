package main

import (
	"betsys/config"
	"betsys/delivery/httpserver"
	"betsys/repository/mysql"
	"betsys/service/serverservice"
)

func main() {
	cfg := config.Config{
		Mysql: mysql.Config{
			Username: "betsys",
			Password: "betsys7",
			Port:     3306,
			Host:     "localhost",
			DBName:   "betsys_db",
		},
		HttpServer: config.HTTPServer{
			Port: 8080,
		},
	}

	serverService := setupServices(cfg)
	server := httpserver.New(cfg, serverService)
	server.Serve()
}

// This should only return services, not sqldb
func setupServices(cfg config.Config) serverservice.Service {
	mySQLrepo := mysql.New(cfg.Mysql)
	serverService := serverservice.New(mySQLrepo)

	return serverService
}
