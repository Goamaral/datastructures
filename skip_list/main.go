package main

import (
	"fmt"
	"math/rand"

	"./skplist"
)

func main() {
	skipList := skplist.New(2, 1)

	for i := 0; i < 10; i += 1 {
		skipList.Insert(rand.Intn(25))
	}

	skipList.PrintLevels()

	var searchItem int
	fmt.Print("Search: ")
	fmt.Scanf("%d", &searchItem)

	if skipList.Search(searchItem) {
		fmt.Println("Item found")
	} else {
		fmt.Println("Item not found")
	}
}
