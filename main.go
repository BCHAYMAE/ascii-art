
package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading banner:", err)
		return
	}

	fmt.Println("Banner loaded successfully!")
	fmt.Println("Number of bytes:", len(data))
}
