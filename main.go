package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args)<2 || len(os.Args)>3{ //check if the number of args is correct
		fmt.Println("Usage: go run . <text> [standard|shadow|thinkertoy]")
		return
	}
	text := os.Args[1] //get the user text
	banner := "standard" //make it by default

	if len(os.Args) == 3 { //when the nbr of args is correct
		banner = os.Args[2] //save the banner the use added if nothing is added by default its standard 
	}
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" { // check whether the banner name is valid
    fmt.Println("Unknown banner:", banner) //error for an unsupported banner
    return 
	}

	data, err := os.ReadFile(banner + ".txt") 
	if err != nil { 
		fmt.Println(err) 
		return 
	}

    lines := strings.Split(string(data), "\n")


		/*
			to have a new line : add\n example:go run . "Hello\nWorld" , 
			but for mint linux:$'...' example:go run . $'Hello\nWorld'
		*/




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