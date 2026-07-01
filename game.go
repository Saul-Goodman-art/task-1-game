package main

import "task-1-gameCases/game/core"

var currentGame *core.Game

func initGame() {
	currentGame = core.NewGame()
}

func handleCommand(command string) string {
	return currentGame.HandleCommand(command)
}
