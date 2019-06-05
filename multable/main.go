package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.Int("start", 2, "Number to start with")
	e := flag.Int("end", 12, "Number to end with")

	flag.Parse()

	for i := *s; i <= *e; i++ {
		for j := *s; j <= *e; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, (i * j))
		}
	}
}
