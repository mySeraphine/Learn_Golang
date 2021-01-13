package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("192.168.1.1:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Println("error!")
			results <- 0
			continue
		}
		conn.Close()
		results <- p

	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var closedport []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i < 1024; i++ {
			ports <- i
		}
	}()
	for i := 1; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		} else {
			closedport = append(closedport, port)
		}
	}
	close(ports)
	close(results)

	sort.Ints(openports)
	sort.Ints(closedport)
	for _, port := range closedport {
		fmt.Printf("%d closed\n", port)
	}
}
