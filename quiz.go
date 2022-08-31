package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	correct, total := 0, 0
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	csv := csv.NewReader(f)
	for {
		ans := 0
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		total++
		fmt.Printf("Question: %s\n", record[0])
		_, err = fmt.Scanln(&ans)
		if err != nil {
			log.Fatal(err)
		}
		corr, err := strconv.Atoi(record[1])
		if corr == ans {
			correct++
			fmt.Println("Correct!")
		}
	}
	fmt.Printf("Correct answers: %d out of %d", correct, total)
}
