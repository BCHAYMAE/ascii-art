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

	var text string 
	fmt.Print("enter a word: ") //display a mssg
	fmt.Scanln(&text) //reads what the user types and store it in char 
	

		for i:=0; i<len(text); i++{
			number := int(text[i]) 	//get the ascii nbr of the current char
			start := 1 + (number-32)*9 //calculate the first line of its drawing 
			//how do we know that !!!!! 
			//cause we are using a standard file wish means we know where each draw start and end
			for j:=start; j<start+8; j++{ //draw the 8 line belongs to that char 
				fmt.Println(lines[j])
			}
  		}
}