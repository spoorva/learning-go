package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvfile, er := os.Open("problems.csv")

	if er != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	csvread := csv.NewReader(csvfile)
	record, er := csvread.ReadAll()

	fmt.Println("Question : " + record[0][0])
	inputReader := bufio.NewReader(os.Stdin)

	text, error := inputReader.ReadString('\n')

	if error != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	correct := 0
	incorrect := 0
	total := 0

	timer := time.After(10 * time.Second)
	for i := 0; i < len(record); i++ {
		total = total + 1
		fmt.Println("Question : " + record[i][0])
		inputReader := bufio.NewReader(os.Stdin)

		text, error = inputReader.ReadString('\n')

		if error != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		if strings.TrimSpace(text) == strings.TrimSpace(record[i][1]) {
			correct = correct + 1
		} else {
			incorrect = incorrect + 1
		}

	}

	select {

	case <-timer:
		fmt.Println("Time's up!")

	}

	fmt.Println("Total number of questions : ")
	fmt.Print(total)
	fmt.Println()
	fmt.Println("Correct responses : ")
	fmt.Print(correct)
	fmt.Println()
	fmt.Println("InCorrect responses : ")
	fmt.Print(incorrect)

}
