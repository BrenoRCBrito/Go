package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i int
	var d float64
	fmt.Scan(&i, &d)
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	fmt.Printf("%d\n%.1f\n%s%s", i+4, d+4, "HackerRank ", s)
}
