package main

import "fmt"

func main() {
	fruits := []string{"apple", "mango", "banana", "joystick"}
oddOneOut := "joystick"
fmt.Println("Which of these is NOT a fruit?")
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}


var answer string
fmt.Print("Type the odd one out: ")
fmt.Scan(&answer)

if answer == oddOneOut {
fmt.Println("Obviously.")
} else if answer != oddOneOut {
fmt.Println("Haha. You're just messing with me, aren't you, with your " + answer + ". It's " + oddOneOut + ", obviously.")
}
} 
