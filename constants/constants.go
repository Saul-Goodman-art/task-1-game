package constants

// Команды
const (
	CmdLook = "осмотреться"
	CmdGo   = "идти"
	CmdWear = "надеть"
	CmdTake = "взять"
	CmdUse  = "применить"
)

// Комнаты / локации
const (
	RoomKitchen  = "kitchen"
	RoomCorridor = "corridor"
	RoomRoom     = "room"
	RoomStreet   = "street"
	Corridor     = "коридор"
	Kitchen      = "кухня"
	Lodging      = "комната"
	Street       = "улица"
)

// Предметы
const (
	ItemBag   = "рюкзак"
	ItemKeys  = "ключи"
	ItemNotes = "конспекты"
	ItemPhone = "телефон"
)

// Цели / объекты
const (
	TargetDoor = "дверь"
)

// Сообщения и ответы (чтобы избежать дублирования строк)
const (
	MsgUnknownCommand        = "неизвестная команда"
	MsgNoPathTo              = "нет пути в "
	MsgNowhereToPut          = "некуда класть"
	MsgNoSuchItem            = "нет такого"
	MsgItemAddedPrefix       = "предмет добавлен в инвентарь: "
	MsgWoreBag               = "вы надели: рюкзак"
	MsgDoorOpened            = "дверь открыта"
	MsgNoItemInInventoryPref = "нет предмета в инвентаре - "
	NothingInteresting       = "ничего интересного. можно пройти - кухня, комната, улица"
	SpringOutside            = "на улице весна. можно пройти - домой"
	EmptyRoom                = "пустая комната. можно пройти - коридор"
	NotUsable                = "не к чему применить"
)
