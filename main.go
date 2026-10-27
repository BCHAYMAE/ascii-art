package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    data, err := os.ReadFile("standard.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    lines := strings.Split(string(data), "\n")

	var char string //a avr can hold the text
	fmt.Print("enter one character: ") //display a mssg
	fmt.Scanln(&char) //reads what the user types and store it in char 
	number := int(char[0])
	fmt.Println("Character number:", number)
	fmt.Println("You entered:", char) //the output

    start := 10

for i := start; i < start+8; i++ {
    fmt.Println(lines[i])
}
}