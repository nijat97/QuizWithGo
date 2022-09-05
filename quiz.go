package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	correct, total := 0, 0
	rand.Seed(time.Now().UnixNano())

	file := flag.String("file", "problems.csv", "file for problems")
	t := flag.Int("time", 30, "timer for game")
	sh := flag.Bool("shuffle", false, "shuffle the questions")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csv := csv.NewReader(f)
	records, err := csv.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if *sh == true {
		rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	}
	fmt.Println("Press Enter to start time.")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(*t) * time.Second)
	go func() {
		<-timer.C
		fmt.Printf("\nCorrect answers: %d out of %d\n", correct, total)
		os.Exit(0)
	}()

	for _, rec := range records {
		ans := ""
		total++
		fmt.Printf("Question: %s\n", rec[0])
		fmt.Scanln(&ans)

		ans = strings.TrimSpace(ans)
		if ans == rec[1] {
			correct++
			fmt.Println("Correct!")
		}
	}
	fmt.Printf("\nCorrect answers: %d out of %d\n", correct, total)
}
