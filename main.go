package main

import (
	"flag"
	"sync"
)

var (
	flgStart  bool
	flgSecure bool
	flgStatic string
	flgHost   string

	serviceWaitGroup sync.WaitGroup
)

func init() {
	serviceWaitGroup = sync.WaitGroup{}
	readArgs()
}

func main() {

	if flgStart {

		serviceWaitGroup.Add(1)

		go serveSite()
		serviceWaitGroup.Wait()
	}

}

func readArgs() {
	flag.BoolVar(&flgStart, "start", false, "to start the web server")
	flag.BoolVar(&flgSecure, "secure", false, "whether to run TLS")
	flag.StringVar(&flgHost, "host", "", "hostname")
	flag.StringVar(&flgStatic, "static", "static", "static files directory")
	flag.Parse()
}
