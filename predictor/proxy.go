package predictor

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Proxy struct {
	Connection net.Conn
}

func NewPredictorProxy() *Proxy {
	newProxy := new(Proxy)
	newProxy.connect()

	return newProxy
}

func (this *Proxy) connect() {
	conn, err := net.Dial("tcp", "localhost:4010")
	connectionError(err, "Error while trying to connect to the predictive model")

	this.Connection = conn
}

func connectionError(err error, message string) {
	if err != nil {
		fmt.Println("Connection Error: ", message, "-> ", err.Error())
		os.Exit(1)
	}
}

func (this *Proxy) Predict(tweet string) bool {
	//this.Connection.Write([]byte(tweet))
	fmt.Fprintf(this.Connection, tweet + "\n")
	//var response []byte
	//length, err := this.Connection.Read(response)
	//fmt.Println(string(response))

	response, err := bufio.NewReader(this.Connection).ReadString('\n')
	connectionError(err, "Error while reading response")
	response = strings.TrimSpace(response)

	var result bool
	if response == "True" {
		result = true
	} else {
		result = false
	}

	return result
}

func (this *Proxy) Close() {
	this.Connection.Close()
}