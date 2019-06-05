package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Rand Seeded successfully.")
}

func main() {
	var nums [3]int
	var feedback string
	nums[0] = rand.Intn(9)
	nums[1] = rand.Intn(9)
	nums[2] = rand.Intn(9)

	fmt.Println("Press enter to proceed")
	b := bufio.NewScanner(os.Stdin)

	if b.Scan() {
		_ = b.Text() //this is just to make it block
	}

	c := 0
	res := ""
	for _, v := range nums {
		if v == 7 {
			c++
		}

		res += strconv.Itoa(v)
	}

	switch c {
	case 0:
		feedback = fmt.Sprintf("Sorry, Seems you are unlucky. Generated number is: %s", res)
		break
	case 1:
		feedback = fmt.Sprintf("Nice luck, 7 came through in %s", res)
		break
	case 2:
		feedback = fmt.Sprintf("Wow!! This is double luck.... congratulations. Generated digit is: %s", res)
		break
	case 3:
		feedback = fmt.Sprintf("Wow! Wow!! Wow!!! You are highly favoured. %s", res)
		break
	default:
		feedback = fmt.Sprintf("Seems something is wrong. This is not recognized")
		break
	}

	fmt.Printf("%s\n", feedback)
}
