package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "text/scanner"
)

func main() {
	name := "hello world"
	age := 4
	zero := "hey"
	fmt.Printf(zero)
	fmt.Println(age)
	fmt.Println(name)
	fmt.Printf("%b", age)
	fmt.Printf("%e", age)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("you typed: %q", input)
}
