package main

import (
	"runtime"
	"flag"
	"log"
	"bufio"
	"fmt"
	"io"
	"os"
	"net"
	"net/http"
	"encoding/json"
)

/*
type location struct {
	lat float64
	long float64
}

type City struct {
	ID string
	Name string `json:"name"`
	Location string `json:"location"`
}
*/

type Config struct {
	Host string `json:host`,
	Port uint16 `json:port`
}

const addr = "127.0.0.1:7070"

func init() {
	log.SetOutput(os.Stdout)
	log.Println("Initializing in runtime ", runtime.Version())
}

func main() {
	log.Println("Parsing command line options")
	shouldCount := flag.Bool("count", false, "In counting mode (query mode if not specified)")
	flag.Parse()
	//dict := make(map[string]int){"one":1, "two":2,"three":3}

	log.Println("Dispatching to actual functionality")
	if *shouldCount == true {
		fmt.Println("Should be counting")

		/*
		s := createServer(addr)
		go s.ListenAndServe()
		*/

		/*
		myCity := City { Name: "Paris", Location: "France"}
		fmt.Printf("The city is %v.\n", myCity)
		bytes, err := json.Marshal(myCity)
		exitOnError(err)
		fmt.Println(string(bytes))
		*/

		//fmt.Println(count(os.Stdin))
	} else {
		conn, _ := net.Dial("tcp", "golang.org:80")
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		status, _ :=  bufio.NewReader(conn).ReadString('\n')
		fmt.Println(status)
	}	

	fmt.Println("Done!")
	os.Exit(0)
}

func exitOnError(err error) {
	if err != nil {
		log.Println("Got an error: ", err)
		os.Exit(1)
	}
}

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server {
		Addr : addr,
		Handler : StringServer("Hello, world!"),
	}
}