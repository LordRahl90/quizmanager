package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var name, feedback string
	var number int
	rand.Seed(time.Now().UnixNano())
	fmt.Println("*****Let's Play a guessing game.*****")

	b := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Your Name pls")

	if b.Scan() {
		name = b.Text()
	}

	fmt.Println("Provide the Number between 1 and 50")
	if b.Scan() {
		n := b.Text()

		for n == "" {
			fmt.Println("Provide the Number between 1 and 50")
			if b.Scan() {
				n = b.Text()
			}
		}

		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("Invalid Number provided")
		}
		number = num
	}

	r := rand.Intn(50)
	if r == number {
		feedback = fmt.Sprintf("Dear %s,\nYou guessed pretty right.\n", name)
	} else {
		feedback = fmt.Sprintf("Dear %s,\nYou guessed: %d\nGenerated Number is: %d\nYou are wrong.\n", name, number, r)
	}

	fmt.Print(feedback)
}
