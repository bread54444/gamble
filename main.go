package main


import (
	"fmt"
	"time"
	"math/rand"
)


func main () {
	//
	
	//random number generator
	rand.Seed(time.Now().UnixNano())
	
	//variables
	currentMoney := 1000
	var bet int 
	var mon int
	var win int
	win = 10000
	//
	fmt.Println("You have a 30% chance to win money")
	fmt.Printf("What would you like the winning number to be, it  must be over $1000, -type [a] for default\n")
	fmt.Scan(&win)

	var isGameOn bool
	isGameOn = true
	if win <= 1000 {
		fmt.Printf("The winning number must be over $1000 \n ")
		fmt.Printf("Restart the program")
		isGameOn = false
	} else {
		fmt.Printf("You have to reach $%d to win the game \n ",win)
	}
	//
	
	
	for {
		mon = rand.Intn(100)
		if isGameOn{
			fmt.Printf("You currently have $%d, How much would you like to bet? \n",currentMoney)
			fmt.Scan(&bet)
		} 
		
	
if bet <= currentMoney && mon >= 70 && isGameOn {
fmt.Printf("You won $%d \n", bet)
currentMoney += bet 
fmt.Printf(" The number was : %d \n",mon)
} else if bet <= currentMoney && mon < 70 && isGameOn {
	fmt.Printf("You lost $%d \n",bet)
	currentMoney -= bet
	fmt.Printf("The number was : %d \n ",mon)
}  else if bet > currentMoney{
	fmt.Println("You can't bet more than you have")
}
//Win game
if currentMoney >= win && win > 1000{
	isGameOn = false
fmt.Println("You have won the game")
break
}

//Lose game

if currentMoney == 0 {
	isGameOn = false
	fmt.Println("You have lost the game")
	break
}

	}


}