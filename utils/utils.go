package utils

import (
	"fmt"
	"strings"
)

func GetStarter() string {
	var player string
	for {
		fmt.Print("Who starts (X or O)? ")
		fmt.Scanf("%s", &player)
		if strings.ToLower(player) == "x" || strings.ToLower(player) == "o" {
			break
		} else {
			fmt.Println("Choose only one of X or O")
		}
	}

	return player
}

func GetInput(player string) (xPos int8, yPos string) {
	for {
		fmt.Printf("[Player %s]\n----------\na, b or c? ", player)
		fmt.Scanf("%s", &yPos)
		if yPos == "a" || yPos == "b" || yPos == "c" {
			break
		} else {
			fmt.Printf("[Player %s] invalid character selected\n", player)
			continue
		}
	}

	for {
		fmt.Printf("[Player %s]\n----------\n1, 2 or 3? ", player)
		fmt.Scanf("%d", &xPos)
		if xPos >= 1 && xPos <= 3 {
			break
		} else {
			fmt.Printf("\n[Player %s] invalid number selected\n", player)
			continue
		}
	}

	return
}

func SwitchPlayer(curr string) string {
	if strings.ToLower(curr) == "x" {
		return "o"
	} else {
		return "x"
	}
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}
