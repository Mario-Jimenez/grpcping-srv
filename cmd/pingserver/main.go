package main

import (
	"fmt"
	"os"

	"github.com/Mario-Jimenez/grpcping-srv/internal/service"
)

func main() {
	srv, err := service.New()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Running server...")
	srv.Run()
}
