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

	reader := bufio.NewReader(os.Stdin) //reader for the input
	fmt.Print("Entrer text: ")

	text, err := reader.ReadString('\n') //read
	
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
		for row:=0; row<8; row++{ //controle drawing rows
			for i:=0; i<len(text); i++{ //controle the char
				number := int(text[i])
				start := 1+(number-32)*9 //calculate where this char's drawing start
				fmt.Print(lines[start+row]) //keeps the char on the same line
			}
			fmt.Println()
		}
}