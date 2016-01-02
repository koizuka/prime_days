package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func isPrime(n int) bool {
	switch {
	case n < 2:
		return false
	case n == 2 || n == 3:
		return true
	case n%2 == 0:
		return false
	}
	end := int(math.Sqrt(float64(n)))
	for i := 3; i <= end; i = i + 2 {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

type IsPrimeResult struct {
	Value       int
	IsPrimeChan chan bool
}

func getPrimeDates(start, end time.Time) []int {
	isPrimeResults := make([]IsPrimeResult, 0)
	for t := start; t.Before(end); t = t.AddDate(0, 0, 1) {
		date, _ := strconv.Atoi(t.Format("20060102"))
		dc := IsPrimeResult{date, make(chan bool)}
		isPrimeResults = append(isPrimeResults, dc)
		go func() {
			dc.IsPrimeChan <- isPrime(dc.Value)
		}()
	}
	dates := make([]int, 0)
	for _, dc := range isPrimeResults {
		if <-dc.IsPrimeChan {
			dates = append(dates, dc.Value)
		}
	}
	return dates
}

func main() {
	year := time.Now().Year()
	if len(os.Args) > 1 {
		var err error
		year, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}

	thisYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	nextYear := thisYear.AddDate(1, 0, 0)
	for _, d := range getPrimeDates(thisYear, nextYear) {
		fmt.Printf("%d\n", d)
	}
}
