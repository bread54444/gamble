
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
	var dif int
	var gCount int
	var gWin int
	var gLoss int
	var winP float64
	var totalG int
	var isGameOn bool
	isGameOn = true
	dif = 70
	win = 10000
	//
	fmt.Printf("What %%chance do you want to win?  \n")
	fmt.Scan(&dif)
	if dif > 99 {
		fmt.Printf("You have to select a number under 100 \n")
		fmt.Printf("Restart the program \n " )
		isGameOn = false
	}
	if isGameOn{
		fmt.Printf("What would you like the winning number to be, it  must be over $1000, -type [a] for default\n")
		fmt.Scan(&win)
	}
	
	
dif = 100 - dif
	
	
	if win <= 1000 {
		fmt.Printf("The winning number must be over $1000 \n ")
		fmt.Printf("Restart the program")
		isGameOn = false
	} else {
		fmt.Printf("You have to reach $%d to win the game \n ",win)
	}
	//
	
	// the game

	for {
		mon = rand.Intn(100)
		//bet system
		if isGameOn{
			fmt.Printf("You currently have $%d, How much would you like to bet? \n",currentMoney)
			fmt.Scan(&bet)
		} 
		
// game thing 
if bet <= currentMoney && mon >= dif && isGameOn {
fmt.Printf(" The number was : %d \n",mon)
fmt.Printf("You won $%d \n", bet)
currentMoney += bet 
gCount++
gWin++
} else if bet <= currentMoney && mon < dif && isGameOn {
	fmt.Printf("The number was : %d \n ",mon)
	fmt.Printf("You lost $%d \n",bet)
	currentMoney -= bet
	gCount++
	gLoss++
}  else if bet > currentMoney{
	fmt.Println("You can't bet more than you have")
}
//Win game
if currentMoney >= win && win > 1000{
	isGameOn = false
	totalG = gLoss + gWin
fmt.Printf("You have won the game with %d wins, %d losses and %d gambles \n",gWin,gLoss,gCount)
winP = float64(gWin) / float64(totalG) * 100
	fmt.Printf("You had a win percentage of %.3f \n",winP)
	
break
}

//Lose game

if currentMoney == 0 {
	totalG = gLoss + gWin
	isGameOn = false
	fmt.Printf("You have lost the game with %d wins, %d losses and %d gambles \n",gWin,gLoss,gCount)
	winP = float64(gWin) / float64(totalG) * 100
	fmt.Printf("You had a win percentage of %.3f \n",winP)
	
	break
}

	}


}
