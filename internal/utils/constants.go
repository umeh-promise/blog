package utils

import "time"

var QueryTimeout = 5 * time.Second
var (
	tokenExp    = time.Hour * 24
	tokenIssuer = "blog"
	authSecret  = GetString("AUTH_SECRET", "basic")
)
