package main

import (
	"runtime"
	"flag"
	"log"
	"bufio"
	"fmt"
	"io"
//	"io/ioutil"
	"os"
//	"net"
	"net/http"
	"time"
//	"encoding/json"
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
	Host string `json:host`
	Port uint16 `json:port`
}

//const addr = "127.0.0.1:7070" // this doesn't work with `http.Get`
const addr = "localhost:7070"

func init() {
	log.SetOutput(os.Stdout)
	log.Println("Initializing in runtime ", runtime.Version())
}

func main() {
	log.Println("Parsing command line options")
	shouldCount := flag.Bool("count", false, "In counting mode (query mode if not specified)")
	//flag.Parse()
	//dict := make(map[string]int){"one":1, "two":2,"three":3}

	/*
	log.Println("Starting the server")

	go func() {
        log.Printf("Serving http at %s\n", addr)
		s := createServer(addr)
		err := s.ListenAndServe()
        if err != nil { panic (err) }
    }()	
	*/
	
	if *shouldCount == true {
		fmt.Println("Should be counting")

		/*
		myCity := City { Name: "Paris", Location: "France"}
		fmt.Printf("The city is %v.\n", myCity)
		bytes, err := json.Marshal(myCity)
		exitOnError(err)
		fmt.Println(string(bytes))
		*/

		//fmt.Println(count(os.Stdin))
	} else {
		/*
		log.Println("Issuing http request to ", addr)
		
		r, err := http.Get("http://" + addr + "/")
		if err != nil {
			log.Println("Got an error: ", err.Error())
			os.Exit(1)
		}
		log.Println("The status is [", r.Status, "]")
		
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		*/

		go fmt.Println("Hello, playground")
		time.Sleep(time.Second)
		/*
		conn, _ := net.Dial("tcp", addr)
		fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost:localhost:7070\r\n\r\n")
		status, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(status)
		*/
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