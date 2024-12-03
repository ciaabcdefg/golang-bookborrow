package main

import (
	"borrow/internal/env"
	"fmt"
)

func main() {
	// envName := os.Args[1]

	envName := ""
	fmt.Scan(&envName)
	fmt.Println(env.LoadAndGet(envName, ""))
}
