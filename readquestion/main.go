package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvFile, err := os.Open("questions.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	l := len(records)

	r := bufio.NewReader(os.Stdin)

	for i := 0; i < l; i++ {
		q := records[i][0]
		ans := records[i][1]
		fmt.Printf("What is the Answer to %s\n", q)
		resp, _ := r.ReadString('\n')
		resp = strings.Trim(resp, "\n")

		if resp == ans {
			fmt.Printf("Very Correct!!!\n")
		} else {
			fmt.Printf("Sadly, You are wrong. The right answer is: %s\n", ans)
		}

		// fmt.Printf("Your Response is: %s, While the answer is: %s\n", strings.Trim(resp, "\n"), ans)
	}
}
