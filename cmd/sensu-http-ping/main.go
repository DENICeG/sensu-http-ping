package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	payload  string
	endpoint string
	method   string
	client   *http.Client
)

func main() {

	app := kingpin.New("sensu-http-ping", "Simple http ping for any received event")
	app.Flag("payload", "Payload for the http ping as string").Default("").Short('p').StringVar(&payload)
	app.Flag("endpoint", "The endpoint which gets called").Required().Short('e').Envar("HTTP_PING_ENDPOINT").StringVar(&endpoint)
	app.Flag("method", "The http method which is used for the ping").Short('m').Envar("HTTP_PING_METHOD").Default("POST").StringVar(&method)
	

	kingpin.MustParse(app.Parse(os.Args[1:]))

	client = new(http.Client)
	req, err := http.NewRequest(method, endpoint, bytes.NewReader([]byte(payload)))
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response:", resp.Status)
}
