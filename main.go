package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

// gq runs query q on endpoint ep and outputs the result.
func gq(q, ep string) error {
	type JSONQuery struct {
		Query string `json:"query"`
	}
	jq, err := json.Marshal(&JSONQuery{Query: q})
	if err != nil {
		return errors.Wrap(err, "marshalling query into JSON")
	}
	resp, err := http.Post(ep, "application/json", bytes.NewReader(jq))
	if err != nil {
		return errors.Wrap(err, "POSTing GraphQL query")
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "reading response from GraphQL server")
	}
	bs2, err := prettify(bs)
	if err != nil {
		return errors.Wrap(err, "prettifying JSON response from GraphQL server")
	}
	if _, err := os.Stdout.Write(bs2); err != nil {
		return errors.Wrap(err, "outputting JSON response")
	}
	return nil
}

func prettify(bs []byte) ([]byte, error) {
	var buf bytes.Buffer
	err := json.Indent(&buf, bs, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, "indenting JSON")
	}
	return buf.Bytes(), nil
}
