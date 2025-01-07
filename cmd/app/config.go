package main

type application struct {
	config baseConfig
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
