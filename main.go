package main

import (
	"bufio"
	"os"
	"strconv"
)

func stdin() int {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	var entry, _ = strconv.Atoi(sc.Text())
	return entry
}

func shareCalculation(entry string) {

}

func main() {
}
