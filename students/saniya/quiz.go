package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	csvFileName := flag.String("csv", "problems.csv", "CSV file with questions and answers")

	// open the file with os.Open which returns a pointer of type os.file
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Error in opening file", err)
	}
	// close the file
	defer file.Close()

	// read the file using NewReader func which takes os.file and returns pointer reader
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Printf("Error while Reading", err)
	}

	correctanswer := 0
	fmt.Println("Press Enter to start the quiz")
	fmt.Scanln()

	// set the timer
	timerlimit := flag.Int("limit", 5, "Timer in seconds")
	timer := time.NewTimer(time.Duration(*timerlimit) * time.Second)
	fmt.Println("Timer started")

	for i, record := range records {

		// Quiz Problem
		fmt.Printf("Problem #%v: %v \n", i+1, record[0])

		answerCh := make(chan string)
		go func() {
			fmt.Printf("Enter your answer\n")
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Timer expired")
			fmt.Printf("You have answered %v correct answers out of %v\n", correctanswer, len(records))
			return
		case answer := <-answerCh:
			{
				if record[1] == answer {
					correctanswer = correctanswer + 1
					fmt.Printf("Correct answer\n")
				} else {
					fmt.Printf("Incorrect answer\n")
				}
			}
		}

	}
	fmt.Printf("You have answered %v correct answers out of %v\n", correctanswer, len(records))
}
