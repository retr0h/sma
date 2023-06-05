package main

import (
	"github.com/retr0h/sma/activity-log/internal/server"
)

func main() {
	println("Starting listening on port 8080")
	srv := server.NewHTTPServer(":8080")
	srv.ListenAndServe()
}
