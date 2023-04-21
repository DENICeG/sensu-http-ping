package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	endpoint       string
	method         string
	payload        string
	insecure       bool
	fail           bool
	timeoutSeconds int
	client         *http.Client
)

func main() {
	start := time.Now()

	app := kingpin.New("sensu-http-ping", "Simple http ping for any received event")
	app.Flag("endpoint", "The endpoint which gets called").Required().Short('e').Envar("HTTP_PING_ENDPOINT").StringVar(&endpoint)
	app.Flag("method", "The http method which is used for the ping").Short('m').Envar("HTTP_PING_METHOD").Default("POST").StringVar(&method)
	app.Flag("payload", "Payload for the http ping as string").Default("").Short('p').StringVar(&payload)
	app.Flag("insecure", "Skip TLS Certificate check").Short('i').BoolVar(&insecure)
	app.Flag("fail", "Exit with code 1 if a non-2xx response has been received").BoolVar(&fail)
	app.Flag("timeout", "Timeout seconds").IntVar(&timeoutSeconds)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	client = new(http.Client)
	if timeoutSeconds > 0 {
		client.Timeout = time.Duration(timeoutSeconds) * time.Second
	}
	if insecure {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecure,
			},
		}
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewReader([]byte(payload)))
	if err != nil {
		fmt.Println("ERR: failed to prepare request:", err.Error(), fmt.Sprintf("[%v]", time.Since(start)))
		os.Exit(2)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERR: request failed:", err.Error(), fmt.Sprintf("[%v]", time.Since(start)))
		os.Exit(2)
	}

	if fail && (resp.StatusCode < 200 || resp.StatusCode >= 300) {
		fmt.Println("ERR: request to", fmt.Sprintf("%q", endpoint), "resulted in code", resp.StatusCode, fmt.Sprintf("[%v]", time.Since(start)))
		os.Exit(2)
	}

	fmt.Println("Response:", resp.Status, fmt.Sprintf("[%v]", time.Since(start)))
}
