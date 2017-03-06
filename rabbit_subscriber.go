package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/brettallred/go-rabbit"
)

var rabbitFlag, exchangeFlag string

func init() {
	flag.Usage = usage
	flag.StringVar(&rabbitFlag, "r", "", "the url of the rabbitmq server")
	flag.StringVar(&exchangeFlag, "e", "events", "the exchange name")
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		usage()
	}

	bindingKey := flag.Arg(0)
	query := flag.Arg(1)

	if rabbitFlag != "" {
		os.Setenv("RABBITMQ_URL", rabbitFlag)
	} else {
		if os.Getenv("RABBITMQ_URL") == "" {
			os.Setenv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
		}
	}

	ch := make(chan []byte, 1)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	cancel := func(sigCh chan os.Signal) chan bool {
		bChan := make(chan bool, 1)
		go func() {
			<-sigCh
			bChan <- true
		}()
		return bChan
	}(sig)

	config := &tapConfig{
		BindingKey: bindingKey,
		Query:      query,
		Exchange:   "events",
		Cancel:     cancel,
		Result:     ch,
	}
	tap(config)

}

type tapConfig struct {
	BindingKey, Query, Exchange string
	Cancel                      <-chan bool
	Result                      chan []byte
}

func tap(config *tapConfig) {
	subscriber := rabbit.Subscriber{
		Concurrency:   1,
		Durable:       false,
		Exchange:      config.Exchange,
		Queue:         "tap",
		RoutingKey:    config.BindingKey,
		PrefetchCount: 100,
		AutoDelete:    true,
	}
	subscriber.PrefixQueueInDev()
	handler := createMessageHandler(config.Query, config.Result)
	rabbit.Register(subscriber, handler)
	rabbit.StartSubscribers()
	select {
	case <-config.Cancel:
		fmt.Println("Stopping tap")
	case result := <-config.Result:
		fmt.Print(string(result))
	}

	// rabbit.CloseSubscribers()
}

func createMessageHandler(query string, ch chan<- []byte) func([]byte) bool {
	b := []byte(query)
	return func(message []byte) bool {
		if bytes.Contains(message, b) {
			ch <- message
		}
		return true
	}
}

func usage() {
	fmt.Printf(usageMessage, os.Args[0])
	os.Exit(1)
}

const usageMessage = `
Usage: %v [-r] [-e] [BINDING_KEY] [QUERY]
The -r flag specifies the url of the rabbitmq server. If flag is not
set, tap will first use any value in the environmental variable
RABBITMQ_URL and finally amqp://guest:guest@localhost:5672/.
The -e flag specifies the topic exchange name. Defaults to "events".