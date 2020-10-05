package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	ep := flag.String("ep", "", "URL of GraphQL endpoint")
	q := flag.String("q", "", "GraphQL query to run")
	flag.Parse()
	if *ep == "" {
		log.Fatal("missing required -ep flag for endpoint")
	}
	if *q == "" {
		log.Fatal("missing required -q flag for GraphQL query")
	}
	if err := gq(*q, *ep); err != nil {
		log.Fatal(err)
	}
}

func gq(q, ep string) error {
	resp, err := http.Post(ep, "application/graphql", strings.NewReader(q))
	if err != nil {
		return errors.Wrap(err, "POSTing GraphQL query")
	}
	defer resp.Body.Close()
	if err != nil {
		return errors.Wrap(err, "reading response from GraphQL server")
	}
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		return errors.Wrap(err, "outputting response")
	}
	return nil
}
