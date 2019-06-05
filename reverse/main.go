package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	b := bufio.NewScanner(os.Stdin)

	fmt.Println("Provide Your Name pls, I wanna print the reverse.")
	if b.Scan() {
		n := b.Text()
		name = n
	}

	l := len(name)
	var rev []rune

	for i := l; i > 0; i-- {
		v := name[i-1]
		rev = append(rev, rune(v))
		// fmt.Printf("%s\n", string(v))
	}

	ans := string(rev)
	fmt.Println(ans)
}
