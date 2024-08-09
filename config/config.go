package config

import "betsys/repository/mysql"

type HTTPServer struct {
	Port int
}

type Config struct {
	Mysql      mysql.Config
	HttpServer HTTPServer
}
