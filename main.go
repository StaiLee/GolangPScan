package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func worker(ports chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range ports {
		address := fmt.Sprintf("test.com:%d", p) // change this to the target address
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func parsePorts(portsStr string) ([]int, error) {
	var ports []int
	portRanges := strings.Split(portsStr, ",")
	for _, pr := range portRanges {
		if strings.Contains(pr, "-") {
			rangeEnds := strings.Split(pr, "-")
			start, err := strconv.Atoi(rangeEnds[0])
			if err != nil {
				return nil, err
			}
			end, err := strconv.Atoi(rangeEnds[1])
			if err != nil {
				return nil, err
			}
			for i := start; i <= end; i++ {
				ports = append(ports, i)
			}
		} else {
			port, err := strconv.Atoi(pr)
			if err != nil {
				return nil, err
			}
			ports = append(ports, port)
		}
	}
	return ports, nil
}

func main() {
	const numWorkers = 100
	const portsToScan = "1-1024" // Change this to specify the ports to scan

	ports := make(chan int, numWorkers)
	results := make(chan int)
	var openPorts []int

	wg := &sync.WaitGroup{}

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ports, results, wg)
	}

	// Parse ports to scan
	portsToScanParsed, err := parsePorts(portsToScan)
	if err != nil {
		fmt.Println("Error parsing ports:", err)
		return
	}

	// Send ports to workers
	go func() {
		for _, port := range portsToScanParsed {
			ports <- port
		}
		close(ports)
	}()

	// Gather results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	for port := range results {
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	// Sort and print open ports
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}
