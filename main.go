package main

import (
	"flag"
	"bufio"
	"fmt"
	"io"
	"os"
	"net"
)

func main() {
	shouldCount := flag.Bool("count", false, "In counting mode (query mode if not specified)")
	flag.Parse()
	//dict := make(map[string]int){"one":1, "two":2,"three":3}
	if *shouldCount == true {
		fmt.Println("Should be counting")
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

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
