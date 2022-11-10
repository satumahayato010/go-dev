package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func stdin() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	entry := sc.Text()
	return entry
}

func someDaysLater(entry string) {
	answer, _ := strconv.Atoi(entry)
	answer = answer * 7
	fmt.Println(answer)
}

func main() {
	si := stdin()
	someDaysLater(si)
}
