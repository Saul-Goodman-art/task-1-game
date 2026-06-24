package main

import "strings"

var place string
var hasBag bool
var hasKeys bool
var hasNotes bool
var keysInRoom bool
var notesInRoom bool
var doorOpen bool

func initGame() {
	place = "kitchen"
	hasBag = false
	hasKeys = false
	hasNotes = false
	keysInRoom = true
	notesInRoom = true
	doorOpen = false
}

func handleCommand(command string) string {
	words := strings.Split(command, " ")

	if words[0] == "осмотреться" {
		return lookAround()
	}

	if words[0] == "идти" && len(words) == 2 {
		return goTo(words[1])
	}

	if words[0] == "надеть" && len(words) == 2 {
		return wear(words[1])
	}

	if words[0] == "взять" && len(words) == 2 {
		return take(words[1])
	}

	if words[0] == "применить" && len(words) == 3 {
		return use(words[1], words[2])
	}

	return "неизвестная команда"
}

func lookAround() string {
	if place == "kitchen" {
		if hasBag {
			return "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"
		}
		return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
	}

	if place == "corridor" {
		return "ничего интересного. можно пройти - кухня, комната, улица"
	}

	if place == "street" {
		return "на улице весна. можно пройти - домой"
	}

	if keysInRoom && notesInRoom && !hasBag {
		return "на столе: ключи, конспекты, на стуле: рюкзак. можно пройти - коридор"
	}
	if keysInRoom && notesInRoom {
		return "на столе: ключи, конспекты. можно пройти - коридор"
	}
	if notesInRoom {
		return "на столе: конспекты. можно пройти - коридор"
	}

	return "пустая комната. можно пройти - коридор"
}

func goTo(newPlace string) string {
	if place == "kitchen" {
		if newPlace == "коридор" {
			place = "corridor"
			return "ничего интересного. можно пройти - кухня, комната, улица"
		}
		return "нет пути в " + newPlace
	}

	if place == "corridor" {
		if newPlace == "кухня" {
			place = "kitchen"
			return "кухня, ничего интересного. можно пройти - коридор"
		}
		if newPlace == "комната" {
			place = "room"
			return "ты в своей комнате. можно пройти - коридор"
		}
		if newPlace == "улица" {
			if doorOpen {
				place = "street"
				return "на улице весна. можно пройти - домой"
			}
			return "дверь закрыта"
		}
	}

	if place == "room" {
		if newPlace == "коридор" {
			place = "corridor"
			return "ничего интересного. можно пройти - кухня, комната, улица"
		}
	}

	return "нет пути в " + newPlace
}

func wear(item string) string {
	if place == "room" && item == "рюкзак" && !hasBag {
		hasBag = true
		return "вы надели: рюкзак"
	}

	return "нет такого"
}

func take(item string) string {
	if !hasBag {
		return "некуда класть"
	}

	if place == "room" && item == "ключи" && keysInRoom {
		keysInRoom = false
		hasKeys = true
		return "предмет добавлен в инвентарь: ключи"
	}

	if place == "room" && item == "конспекты" && notesInRoom {
		notesInRoom = false
		hasNotes = true
		return "предмет добавлен в инвентарь: конспекты"
	}

	return "нет такого"
}

func use(item string, target string) string {
	if item == "ключи" && !hasKeys {
		return "нет предмета в инвентаре - ключи"
	}
	if item == "телефон" {
		return "нет предмета в инвентаре - телефон"
	}

	if item == "ключи" && target == "дверь" {
		doorOpen = true
		return "дверь открыта"
	}

	return "не к чему применить"
}
