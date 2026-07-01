package core

import (
	"task-1-gameCases/constants"
)

type rule struct {
	cond func(*Game, *Room) bool
	text string
}

func (game *Game) createRooms() map[string]*Room {
	return map[string]*Room{
		constants.RoomKitchen:  createRoomKitchen(game),
		constants.RoomCorridor: createRoomCorridor(),
		constants.RoomRoom:     createRoomRoom(game),
		constants.RoomStreet:   createRoomStreet(game),
	}
}

func createRoomKitchen(game *Game) *Room {
	return &Room{
		look: game.lookKitchen,
		exits: map[string]Exit{
			constants.Corridor: {
				to:      constants.RoomCorridor,
				message: constants.NothingInteresting,
			},
		},
	}
}

func createRoomCorridor() *Room {
	return &Room{
		look: func(*Game) string {
			return constants.NothingInteresting
		},
		exits: map[string]Exit{
			constants.Kitchen: {
				to:      constants.RoomKitchen,
				message: "кухня, ничего интересного. можно пройти - коридор",
			},
			constants.Lodging: {
				to:      constants.RoomRoom,
				message: "ты в своей комнате. можно пройти - коридор",
			},
			constants.Street: {
				to:      constants.RoomStreet,
				message: constants.SpringOutside,
				locked: func(game *Game) string {
					if !game.doorOpen {
						return "дверь закрыта"
					}
					return ""
				},
			},
		},
	}
}

func createRoomRoom(game *Game) *Room {
	return &Room{
		look: game.lookRoom,
		exits: map[string]Exit{
			constants.Corridor: {
				to:      constants.RoomCorridor,
				message: constants.NothingInteresting,
			},
		},
		items: map[string]bool{
			constants.ItemKeys:  true,
			constants.ItemNotes: true,
		},
	}
}

func createRoomStreet(game *Game) *Room {
	return &Room{
		look: func(*Game) string {
			return constants.SpringOutside
		},
	}
}

func (game *Game) lookKitchen(*Game) string {
	if game.hasItem(constants.ItemBag) {
		return "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"
	}
	return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
}

func (game *Game) lookRoom(*Game) string {
	room := game.rooms[constants.RoomRoom]
	rules := createRules()

	for _, rl := range rules {
		if rl.cond(game, room) {
			return rl.text
		}
	}

	return constants.EmptyRoom
}

func createRules() []rule {
	return []rule{
		{
			// ключи и конспекты есть, рюкзака нет в инвентаре
			cond: func(g *Game, r *Room) bool {
				return r.items[constants.ItemKeys] && r.items[constants.ItemNotes] && !g.hasItem(constants.ItemBag)
			},
			text: "на столе: ключи, конспекты, на стуле: рюкзак. можно пройти - коридор",
		},
		{
			// ключи и конспекты есть (рюкзак может быть надет)
			cond: func(g *Game, r *Room) bool {
				return r.items[constants.ItemKeys] && r.items[constants.ItemNotes]
			},
			text: "на столе: ключи, конспекты. можно пройти - коридор",
		},
		{
			// только конспекты
			cond: func(g *Game, r *Room) bool {
				return r.items[constants.ItemNotes]
			},
			text: "на столе: конспекты. можно пройти - коридор",
		},
	}
}
