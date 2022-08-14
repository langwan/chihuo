package main

import (
	"bufio"
	"fmt"
	"os"
)

func worker(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("worker id = %d\n", id)
	}
}

func main() {

	go worker(1)
	go worker(2)

	buf := bufio.NewReader(os.Stdin)
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(sentence))
	}

}
