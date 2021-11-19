package main

import (
	"bufio"
	"fmt"
	"os"

	"./hshlist"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Source data: ")
	sourcedata, _ := reader.ReadSlice('\n')
	sourcedata = sourcedata[:len(sourcedata)-1]
	sourceHashlist := hshlist.New(sourcedata)

	fmt.Print("Target data: ")
	targetData, _ := reader.ReadSlice('\n')
	targetData = targetData[:len(targetData)-1]
	targetHashlist := hshlist.New(targetData)

	if hshlist.Match(sourceHashlist, targetHashlist) {
		fmt.Println("Data is the same")
	} else {
		fmt.Println("Data is different")
	}
}
