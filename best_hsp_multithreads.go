package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var (
	blast         string
	max_evalue    float64
	min_bit_score float64
	err           error
)

func init() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Usage: best_hsp <alignment> <evalue> <bit_score>")
		os.Exit(1)
	}

	blast = os.Args[1]

	max_evalue, err = strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "evalue should be float")
		os.Exit(1)
	}

	min_bit_score, err = strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bit_score should be float")
		os.Exit(1)
	}

}

func check_one_record(line string, fields []string, out chan string) {
	evalue, err := strconv.ParseFloat(strings.Trim(fields[10], " "), 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to parse evalue (%s) of %s\n", fields[10], fields[0])
		os.Exit(1)
	}
	bit_score, err := strconv.ParseFloat(strings.Trim(fields[11], " "), 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to parse bit_score (%s) of %s\n", fields[11], fields[0])
		os.Exit(1)
	}

	// fmt.Println(fields[0], max_evalue, min_bit_score, evalue, bit_score)
	if evalue <= max_evalue && bit_score >= min_bit_score {
		out <- line
	}
}

func main() {
	fh, err := os.Open(blast)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to open file: %s\n", blast)
		os.Exit(1)
	}

	defer func() {
		fh.Close()
	}()

	threads := runtime.NumCPU()
	runtime.GOMAXPROCS(threads)

	tokens := make(chan int, threads)
	var wg sync.WaitGroup

	// output channel
	out := make(chan string)
	go func() {
		for {
			select {
			case str := <-out:
				fmt.Println(str)
			}
		}
	}()

	reader := bufio.NewReader(fh)
	var line string
	var query string = ""
	for {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if line[0] == '#' {
			continue
		}
		line = strings.TrimRight(line, "\r\n")
		fields := strings.Split(line, "\t")

		if query != fields[0] {
			tokens <- 1
			wg.Add(1)

			go func(line string, fields []string, out chan string) {
				defer func() {
					wg.Done()
					<-tokens
				}()
				check_one_record(line, fields, out)
			}(line, fields, out)

			query = fields[0]
		}
	}

	wg.Wait()
}
