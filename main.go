package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"flag"
	"time"
)

type problem struct {
	question string
	answer string
}

func parse_lines(lines [][]string) []problem {
	res := make([]problem, len(lines))
	for i, line := range lines {
		res[i] = problem{
			question: line[0],
			answer: line[1], //strings.TrimSpace(line[1])
		}
	}
	return res
}

func main() {
	csv_file := flag.String("csv", "problems.csv", "Enter a csv file in the format of 'question,answer'")
	time_limit := flag.Int("timelimit", 30, "Answer the question within the specified number of seconds")
	flag.Parse()
	csv_fd, err := os.Open(*csv_file)
	if err != nil {
		csv_fd.Close()
		fmt.Printf("Failed to open file %s\n", *csv_file)
		log.Fatal(err)
	}

	file_reader := csv.NewReader(csv_fd)
	lines, err := file_reader.ReadAll()
	if err != nil {
		csv_fd.Close()
		log.Fatal(err)
	}

	problems := parse_lines(lines)
	score := 0
	timer := time.NewTimer(time.Duration(*time_limit) * time.Second)
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, prob.question)
		respCh := make(chan string)
		go func() {
			var response string
			fmt.Scanf("%s\n", &response)
			respCh <- response
		}()

		select {
		case <-timer.C:
			fmt.Printf("Time's up! Your score is %d out of %d.\n", score, len(problems))
			return
		case response := <-respCh:
			if response == prob.answer {
				fmt.Println("Correct :)")
				score += 1
			} else {
				fmt.Printf("Wrong, answer is %s :(\n", prob.answer)
			}
		}
	}
	fmt.Printf("Your score is %d out of %d.\n", score, len(problems))
	
	csv_fd.Close()
}