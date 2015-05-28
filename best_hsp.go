package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func main() {
	fh, err := os.Open(blast)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to open file: %s\n", blast)
		os.Exit(1)
	}

	defer func() {
		fh.Close()
	}()

	reader := bufio.NewReader(fh)

	var line string
	query := ""
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
				fmt.Println(line)
			}
			query = fields[0]
		}
	}
}
