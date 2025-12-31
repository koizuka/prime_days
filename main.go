package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
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

func groupByMonth(dates []int) [][]int {
	months := make([][]int, 12)
	for i := range months {
		months[i] = []int{}
	}
	for _, d := range dates {
		month := (d / 100) % 100
		if month >= 1 && month <= 12 {
			months[month-1] = append(months[month-1], d)
		}
	}
	return months
}

func main() {
	group := flag.Bool("group", false, "group output by month (one line per month)")
	flag.Parse()

	year := time.Now().Year()
	if flag.NArg() > 0 {
		var err error
		year, err = strconv.Atoi(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	}

	thisYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	nextYear := thisYear.AddDate(1, 0, 0)
	dates := getPrimeDates(thisYear, nextYear)

	if *group {
		fmt.Printf("%d prime dates in %d:\n", len(dates), year)
		for _, monthDates := range groupByMonth(dates) {
			strs := make([]string, len(monthDates))
			for i, d := range monthDates {
				strs[i] = strconv.Itoa(d)
			}
			fmt.Println(strings.Join(strs, " "))
		}
	} else {
		for _, d := range dates {
			fmt.Printf("%d\n", d)
		}
	}
}
