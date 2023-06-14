package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	csv_fd, err := os.Open("problems.csv")
	if err != nil {
		csv_fd.Close()
		log.Fatal(err)
	}
	fmt.Println("CSV file opened successfully")

	file_reader := csv.NewReader(csv_fd)
	questions, err := file_reader.ReadAll()
	if err != nil {
		csv_fd.Close()
		log.Fatal(err)
	}
	fmt.Println("CSV file read successfully")
	fmt.Println(questions)

	var score int = 0
	for _, row := range questions {
		question := row[0]
		answer := row[1]
		fmt.Printf("What is %s?\n", question)
		var response string
		fmt.Scanln(&response)
		if response == answer {
			score += 1
		}
	}
	fmt.Printf("Your score is %s out of %s.\n", strconv.Itoa(score), strconv.Itoa(len(questions)))
	
	csv_fd.Close()
}