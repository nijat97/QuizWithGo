package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	correct, total := 0, 0
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	csv := csv.NewReader(f)
	fmt.Println("Press Enter to start time.")
	fmt.Scanln()
	timer := time.NewTimer(10 * time.Second)
	go func() {
		<-timer.C
		fmt.Printf("Correct answers: %d out of %d\n", correct, total)
		os.Exit(0)
	}()
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
		fmt.Scanln(&ans)
		corr, err := strconv.Atoi(record[1])
		if corr == ans {
			correct++
			fmt.Println("Correct!")
		}
	}
}
