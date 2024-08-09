package main

import (
	"betsys/config"
	"betsys/delivery/httpserver"
	"betsys/repository/mysql"
)

func main() {
	cfg := config.Config{
		Mysql: mysql.Config{
			Username: "question",
			Password: "password",
			Host:     "localhost",
			Port:     3306,
		},
		HttpServer: config.HTTPServer{
			Port: 8080,
		},
	}

	//serverService := setupServices(cfg)
	//server := httpserver.New(cfg, serverService)
	server := httpserver.New(cfg)
	server.Serve()
}

// This should only return services, not sqldb
//func setupServices(cfg config.Config) *mysql.MySQLDB {
//	mySQLrepo := mysql.New(cfg.Mysql)
//
//	return mySQLrepo
//}
