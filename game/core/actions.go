package core

import (
	consts "task-1-gameCases/constants"
)

func (game *Game) lookAround(args []string) string {
	if len(args) != 0 {
		return consts.MsgUnknownCommand
	}
	return game.currentRoom().look(game)
}

func (game *Game) goTo(args []string) string {
	if len(args) != 1 {
		return consts.MsgUnknownCommand
	}

	destination := args[0]
	exit, ok := game.currentRoom().exits[destination]
	if !ok {
		return consts.MsgNoPathTo + destination
	}

	if exit.locked != nil {
		message := exit.locked(game)
		if message != "" {
			return message
		}
	}

	game.place = exit.to
	return exit.message
}

func (game *Game) wear(args []string) string {
	if len(args) != 1 {
		return consts.MsgUnknownCommand
	}

	item := args[0]
	props, ok := game.itemProps[item]
	if !ok || !props.Wearable {
		return consts.MsgNoSuchItem
	}

	if item == consts.ItemBag && game.place == consts.RoomRoom && !game.hasItem(consts.ItemBag) {
		game.addItem(consts.ItemBag)
		return consts.MsgWoreBag
	}

	return consts.MsgNoSuchItem
}

func (game *Game) take(args []string) string {
	if len(args) != 1 {
		return consts.MsgUnknownCommand
	}

	if !game.hasItem(consts.ItemBag) {
		return consts.MsgNowhereToPut
	}

	item := args[0]
	currentRoom := game.currentRoom()
	if currentRoom.items == nil || !currentRoom.items[item] {
		return consts.MsgNoSuchItem
	}

	currentRoom.items[item] = false
	game.addItem(item)
	return consts.MsgItemAddedPrefix + item
}

func (game *Game) use(args []string) string {
	if len(args) != 2 {
		return consts.MsgUnknownCommand
	}

	item, target := args[0], args[1]

	if !game.hasItem(item) {
		return consts.MsgNoItemInInventoryPref + item
	}

	props, ok := game.itemProps[item]
	if !ok {
		return consts.NotUsable
	}

	if handler, ok := props.UsableOn[target]; ok && handler != nil {
		return handler(game)
	}

	return consts.NotUsable
}
