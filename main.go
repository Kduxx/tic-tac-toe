package main

import (
	"fmt"
	"strings"

	"github.com/Kduxx/tic-tac-toe/game"
	"github.com/Kduxx/tic-tac-toe/utils"
)

func main() {
	utils.Clear()
	tick := game.TickTackToe{}
	player := utils.GetStarter()
	utils.Clear()
	tick.PrintGame()
	for {
		xPos, yPos := utils.GetInput(player)
		err := tick.AddPlay(player, xPos-1, yPos)
		if err != nil {
			if strings.Contains(err.Error(), "[PositionInUse]") {
				utils.Clear()
				fmt.Println("\nThe play is invalid!")
			} else {
				utils.Clear()
				fmt.Printf("Unexpected error: %e\n", err)
			}
		} else {
			win, won, draw := tick.CheckGame()
			if win {
				utils.Clear()
				tick.PrintGame()
				fmt.Printf("Player %s Won!\n\n", won)
				break
			} else if draw {
				utils.Clear()
				tick.PrintGame()
				fmt.Println("Game ended in draw")
				break
			}
			// Change Player
			player = utils.SwitchPlayer(player)
			utils.Clear()
		}

		tick.PrintGame()
	}
}
