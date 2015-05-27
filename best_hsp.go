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
		evalue, _ := strconv.ParseFloat(fields[10], 64)
		bit_score, _ := strconv.ParseFloat(fields[11], 64)

		if query != fields[0] {
			if evalue <= max_evalue && bit_score >= min_bit_score {
				fmt.Println(line)
			}
			query = fields[0]
		}
	}
}
