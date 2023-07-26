package main

import (
	"fmt"
	"log"

	"github.com/stiorg/mightydragon/api"
)

func main() {
	svc := api.NewCatFactService("https://catfact.ninja/fact")
	svc = api.NewLoggingService(svc)

	apiServer := api.NewApiServer(svc)

	port := "3000"

	fmt.Printf("starting catfact service at port %s\n", port)
	log.Fatal(apiServer.Start(":" + port))
}
