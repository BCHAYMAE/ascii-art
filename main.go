package main

import (
	"bufio"
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

	reader := bufio.NewReader(os.Stdin) // input terminal
	fmt.Print("Enter text: ")     

	var input strings.Builder //to collect all input lines.
	for { // keep reading lines until the user enters the end-of-input signal

		/*
		to send end of input signal in terminal
		-windows:press Ctrl + Z then Enter
		-linix mint: press Ctrl + D then Enter
		*/

		line, err := reader.ReadString('\n') // Read one line, including its newline
		input.WriteString(line) // add the line to the collected input
		if err != nil { 
			break 
		}
	}
	text := input.String() // convert the input into one string
	text = strings.TrimSuffix(text, "\n") 
	text = strings.TrimSuffix(text, "\r") 
	
	if err != nil && len(text) == 0 { 
		fmt.Println(err) 
		return            
	}

	text = strings.TrimSuffix(text, "\n") 
	text = strings.TrimSuffix(text, "\r")

		/*for i:=0; i<len(text); i++{
			number := int(text[i]) 	//get the ascii nbr of the current char
			start := 1 + (number-32)*9 //calculate the first line of its drawing 
			//how do we know that !!!!! 
			//cause we are using a standard file wish means we know where each draw start and end
			for j:=start; j<start+8; j++{ //draw the 8 line belongs to that char 
				fmt.Println(lines[j])
			}
  		}*/

		//instead we need to print line by line of each char in the word
		
for _, word := range strings.Split(text, "\n") { // split the input wherever a newline occurs
    for row := 0; row < 8; row++ { // print the word one drawing row at a time
        for i := 0; i < len(word); i++ { 
            number := int(word[i]) // get the ASCII nbr 
            if number < 32 || number > 126 { // check whether the character is printable ASCII
                fmt.Print("?") //for unsupported characters
                continue // Skip to the next character
            }

            start := 1 + (number-32)*9 //calculate where the drawing starts in the file

            if start+row < len(lines) { // check that the line exists in the file
                fmt.Print(lines[start+row]) // print this row of the character
            }
        }

        fmt.Println() // move to the next terminal line after printing the row
    }
}

}