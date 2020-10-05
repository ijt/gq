package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var ep = flag.String("ep", "", "URL of GraphQL endpoint")
var q = flag.String("q", "", "GraphQL query to run")

func main() {
	flag.Parse()
	if *ep == "" {
		log.Fatal("missing required -ep flag for endpoint")
	}
	if *q == "" {
		log.Fatal("missing required -q flag for GraphQL query")
	}
	resp, err := http.Post(*ep, "application/graphql", strings.NewReader(*q))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
}
