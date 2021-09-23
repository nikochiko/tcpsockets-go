package loadtest

import (
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"time"

	"github.com/nikochiko/tcpsockets-go/util"
)

const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

// LoadTest load tests a server with X concurrent TCP connections
func LoadTest(service string, concurrency int, writeData []byte) error {
	if concurrency < 0 {
		err := fmt.Errorf("concurrency must be greater than 0")
		util.CheckError(err)
	}

	wg := sync.WaitGroup{}
	responseTimes := make([]int, concurrency)

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go connectTCP(&wg, responseTimes, i, service, writeData)
	}

	fmt.Println("waiting for load tests to finish")
	wg.Wait()
	endTime := time.Now()

	fmt.Printf("Total time taken:\t %d ms\n", endTime.Sub(startTime).Milliseconds())
	fmt.Printf("Max wait time:   \t %d ms\n", Max(responseTimes))
	fmt.Printf("Min wait time:   \t %d ms\n", Min(responseTimes))
	fmt.Printf("Mean wait time:  \t %g ms\n", Mean(responseTimes))
	fmt.Printf("Median wait time:\t %g ms\n", Median(responseTimes))

	return nil
}

func connectTCP(wg *sync.WaitGroup, responseTimes []int, id int, service string, writeData []byte) {
	defer wg.Done()

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	util.CheckError(err)

	startTime := time.Now()

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	util.CheckError(err)

	defer conn.Close()

	conn.Write(writeData)

	returnedData, err := ioutil.ReadAll(conn)
	util.CheckError(err)

	endTime := time.Now()

	responseTime := endTime.Sub(startTime).Milliseconds()
	responseTimes[id] = int(responseTime)

	fmt.Printf("[%d ms] Got output: %s\n", endTime.Sub(startTime).Milliseconds(), string(returnedData))
}

func Mean(array []int) float64 {
	sum := 0
	for _, num := range array {
		sum += num
	}

	return float64(sum) / float64(len(array))
}

func Median(array []int) (median float64) {
	if len(array) % 2 == 0 {
		middleIndex1 := len(array) / 2
		middleIndex2 := middleIndex1 + 1

		median = float64(array[middleIndex1] + array[middleIndex2]) / 2.0
	} else {
		middleIndex := len(array) / 2

		median = float64(array[middleIndex])
	}

	return
}

func Max(array []int) int {
	max := MinInt

	for _, n := range array {
		if n > max {
			max = n
		}
	}

	return max
}

func Min(array []int) int {
	min := MaxInt

	for _, n := range array {
		if n < min {
			min = n
		}
	}

	return min
}
