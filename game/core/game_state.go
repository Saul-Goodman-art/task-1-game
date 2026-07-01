package core

import (
	"strings"
	"task-1-gameCases/constants"
)

type Game struct {
	place     string
	inventory map[string]bool
	rooms     map[string]*Room
	actions   map[string]Action
	doorOpen  bool
	itemProps map[string]ItemProps
}

type Room struct {
	look  func(*Game) string
	exits map[string]Exit
	items map[string]bool
}

type Exit struct {
	to      string
	message string
	locked  func(*Game) string
}

type Action func(args []string) string

type ItemProps struct {
	Wearable bool
	UsableOn map[string]func(*Game) string
}

func NewGame() *Game {
	game := &Game{
		place:     constants.RoomKitchen,
		inventory: map[string]bool{},
		doorOpen:  false,
		itemProps: map[string]ItemProps{},
	}
	game.initGame()
	return game
}

func (game *Game) initGame() {
	game.toInitItemMetadata()
	game.rooms = game.createRooms()
	game.toBindCommandsActions()
}

func (game *Game) toBindCommandsActions() {
	game.actions = map[string]Action{
		constants.CmdLook: game.lookAround,
		constants.CmdGo:   game.goTo,
		constants.CmdWear: game.wear,
		constants.CmdTake: game.take,
		constants.CmdUse:  game.use,
	}
}

func (game *Game) toInitItemMetadata() {

	game.itemProps[constants.ItemBag] = ItemProps{
		Wearable: true,
		UsableOn: map[string]func(*Game) string{},
	}
	game.itemProps[constants.ItemKeys] = ItemProps{
		Wearable: false,
		UsableOn: map[string]func(*Game) string{
			constants.TargetDoor: func(g *Game) string {
				g.doorOpen = true
				return constants.MsgDoorOpened
			},
		},
	}
}

func (game *Game) HandleCommand(input string) string {
	parts := strings.Fields(strings.TrimSpace(input))
	if len(parts) == 0 {
		return constants.MsgUnknownCommand
	}

	action, ok := game.actions[parts[0]]
	if !ok {
		return constants.MsgUnknownCommand
	}

	return action(parts[1:])
}

func (game *Game) currentRoom() *Room {
	return game.rooms[game.place]
}

func (game *Game) hasItem(item string) bool {
	return game.inventory[item]
}

func (game *Game) addItem(item string) {
	game.inventory[item] = true
}
