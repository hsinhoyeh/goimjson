package main

import (
	"flag"
	"fmt"
	"log"

	imjHTTP "github.com/hsinhoyeh/goimjson/http"
)

var (
	port = flag.Int("port", 9000, "the service port for imjson")
)

func main() {
	log.Print("listen :%v\n", *port)
	log.Fatal(imjHTTP.ListenAndServe(fmt.Sprintf(":%v", *port)))
}
