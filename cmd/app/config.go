package main

import "go.uber.org/zap"

type application struct {
	config baseConfig
	logger *zap.SugaredLogger
}

type baseConfig struct {
	addr string
	env  string
	db   dbConfig
}

type dbConfig struct {
	addr        string
	maxOpenConn int
	maxIdleConn int
	maxIdleTime string
}
