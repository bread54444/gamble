package main


import (
	"fmt"
	"time"
	"math/rand"
	"errors"
	"os"
	"io/ioutil"
	"bufio"
	
)

func writeFile(contents string) (*os.File){
	var endFile, err = os.OpenFile("Stats.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0600)
	defer endFile.Close()

	if err != nil {
		panic("Error creating file ")
	}
	writer := bufio.NewWriter(endFile)
	fmt.Fprintln(writer,contents)
	writer.Flush()
	readFile,err := ioutil.ReadFile("Stats.txt")
	fmt.Print(readFile)
	if err != nil{
		panic("File not found")
	}
	return endFile
}

func calcWinP(w, g int) float64 {
	return float64(w) / float64(g) * 100

}
type Settings struct {
	dif int
	win int
}
func main () {
	// func
	
	//random number generator
	rand.Seed(time.Now().UnixNano())
	var(
		 currentMoney int
		 bet int 
		 mon int
		 oldDif int
		 gCount int
		 gWin int
		 gLoss int
		//var winP float64
		 totalG int
		 totalMon int
		 isGameOn bool
		 settings Settings
		 setCheck int
		 winPer float64
		// fileContents []byte
		 val string
	)
	currentMoney = 1000
	isGameOn = true
	
	//
	fmt.Printf("Enter 1 for default settings(Reccomended for best user expierence), enter 2 for customized settings \n")
	fmt.Scan(&setCheck)
	if setCheck == 1 {
		settings = Settings{dif:50,win:10000}

	} else if setCheck == 2 {

	} else {
		err := errors.New("Error : use 1 or 2 \n")
		fmt.Print(err)
		os.Exit(1)
	}
	if setCheck != 1 {
	fmt.Printf("What %%chance do you want to win?  \n")
	fmt.Scan(&settings.dif)
	}
	if settings.dif > 99 {
		err := errors.New(" Error : The percentage must be under 99 \n")
		fmt.Print(err)
		
		isGameOn = false
		os.Exit(1)
	}
	if isGameOn && setCheck != 1{
		fmt.Printf("What would you like the winning number to be, it  must be over $1000\n")
		fmt.Scan(&settings.win)
	}
	
oldDif = settings.dif
settings.dif = 100 - settings.dif
	
	
	if settings.win <= 1000 {
		err:= errors.New("Error : The winning number must be over $1000 \n")
		fmt.Print(err)
		os.Exit(1)
		isGameOn = false
	} else {
		fmt.Printf("You have to reach $%d to win the game \n ",settings.win)
		fmt.Printf("You have a %v percent chance to win, The randomized number has to be over %v for you to win the bet \n",oldDif,settings.dif)
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
if bet <= currentMoney && mon >= settings.dif && isGameOn {
fmt.Printf("The random number was : %d \n",mon)
fmt.Printf("You won $%d \n", bet)
currentMoney += bet 
gCount++
gWin++
totalMon = totalMon + bet
} else if bet <= currentMoney && mon < settings.dif && isGameOn {
	fmt.Printf("The number was : %d \n ",mon)
	fmt.Printf("You lost $%d \n",bet)
	currentMoney -= bet
	gCount++
	gLoss++
	totalMon = totalMon + bet

}  else if bet > currentMoney{
	err:= errors.New("Error: You can't bet more than you have \n")
	fmt.Print(err)
} 
//Win game
if currentMoney >= settings.win && settings.win > 1000{
	isGameOn = false
	totalG = gLoss + gWin
fmt.Printf("You have won the game with %d wins, %d losses and %d gambles \n",gWin,gLoss,gCount)
fmt.Printf("You bet a total of $%d dollars", totalMon)
fmt.Println(calcWinP(gWin,totalG))

	
break
}

//Lose game

if currentMoney == 0 {
	totalG = gLoss + gWin
	isGameOn = false

p1 := fmt.Sprintf("You have lost the game with %d wins, %d losses and %d gambles \n",gWin,gLoss,gCount)
	winPer = calcWinP(gWin,totalG)
p2 := fmt.Sprintf("You had a win percentage of %v%% \n",winPer)
p3 := fmt.Sprintf("You have bet a total of $%d dollars", totalMon)
barrier := fmt.Sprintf("------------------------------------------- %v",val)
endString := fmt.Sprintf("%v \n %v \n %v \n %v \n",p1,p2,p3,barrier)
//fileContents = []byte(bigString)
writeFile(endString)
	break
}

	}


}

