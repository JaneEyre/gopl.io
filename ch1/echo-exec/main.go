// Echo exercise
// 1) os.Args[0]
// 2) print both index and value

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	for i, arg := range os.Args[1:] {
		fmt.Println(i, " + ", arg)
	}
}
