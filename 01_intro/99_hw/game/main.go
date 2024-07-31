package main

import (
	"strings"

	"golang.org/x/exp/slices"
)

type Room struct {
	name        string
	description []string //ключ для осмотреться, значение для идти
	exits       map[string]string
	items       []string
}
type Player struct {
	currentRoom *Room
	inventory   []string
}

var rooms map[string]*Room
var player Player

func initGame() {
	rooms = make(map[string]*Room)

	rooms["кухня"] = &Room{
		name:        "кухня",
		description: []string{"ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор", "кухня, ничего интересного. можно пройти - коридор"},
		exits:       map[string]string{"коридор": "коридор"},
	}

	rooms["коридор"] = &Room{
		name:        "коридор",
		description: []string{"", "ничего интересного. можно пройти - кухня, комната, улица"},
		exits:       map[string]string{"кухня": "кухня", "комната": "комната", "улица": "улица"},
	}

	rooms["комната"] = &Room{
		name:        "комната",
		description: []string{"на столе: ключи, конспекты, на стуле: рюкзак. можно пройти - коридор", "ты в своей комнате. можно пройти - коридор"},
		exits:       map[string]string{"коридор": "коридор"},
		items:       []string{"ключи", "конспекты", "рюкзак"},
	}
	rooms["улица"] = &Room{
		name:        "улица",
		description: []string{"", "дверь закрыта"},
		exits:       map[string]string{"кухня": "кухня"},
	}

	player = Player{
		currentRoom: rooms["кухня"],
		inventory:   []string{}}
}

/*func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}*/

func (p *Player) lookAround() string {
	if p == nil {
		return "Игрок не найден"
	}
	if p.currentRoom == nil {
		return "комната не найдена"
	}
	if len(p.currentRoom.description) == 0 {
		return "Нет описания комнаты"
	}
	return p.currentRoom.description[0]
}

func (p *Player) putOn(item string) string {
	p.inventory = append(p.inventory, item)
	rooms["комната"].description[0] = "на столе: ключи, конспекты. можно пройти - коридор"
	player.currentRoom.changeItems("рюкзак")
	return "вы надели: " + item
}

func (p *Player) use(how, where string) string {
	if !slices.Contains(p.inventory, how) {
		return "нет предмета в инвентаре - " + how
	}
	if how == "ключи" && where == "дверь" {
		rooms["улица"].changeDescription1("на улице весна. можно пройти - домой")
		//p.currentRoom.description[1] = "на улице весна. можно пройти - домой"
		return "дверь открыта"
	}
	return "не к чему применить"
}

func (p *Player) move(direction string) string { // создали метод идти
	if p == nil || p.currentRoom == nil {
		return "Комнаты или игрока не сущетсвует"
	}
	curRoomName := p.currentRoom.name
	nextRoomName := p.currentRoom.exits[direction]
	if nextRoomName == "" {
		return "нет пути в " + direction
	}

	p.currentRoom = rooms[nextRoomName]
	if p.currentRoom.description[1] == "дверь закрыта" {
		p.currentRoom = rooms[curRoomName]
		return "дверь закрыта"

	}
	return p.currentRoom.description[1]
}

func (r *Room) changeDescription(newDesc string) {
	r.description[0] = newDesc
}
func (r *Room) changeDescription1(newDesc string) {
	r.description[1] = newDesc
}
func (r *Room) changeItems(item string) {
	for i := range r.items {
		if r.items[i] == item {
			r.items[i] = ""
			break
		}
	}

}

func (p *Player) take(item string) string {
	if !slices.Contains(p.inventory, "рюкзак") {
		return "некуда класть"
	}
	if strings.Contains(p.currentRoom.description[0], item) {
		p.inventory = append(p.inventory, item)
		p.currentRoom.changeDescription(strings.Replace(p.currentRoom.description[0], item+", ", "", 1))
		p.currentRoom.changeItems(item)
		if p.currentRoom.items[0] == "" && p.currentRoom.items[1] == "" && p.currentRoom.items[2] == "" {
			p.currentRoom.changeDescription("пустая комната. можно пройти - коридор")
		}
		return "предмет добавлен в инвентарь: " + item
	} else {
		return "нет такого"
	}

}

func handleCommand(command string) string {
	parts := strings.Split(command, " ")
	switch parts[0] {
	case "осмотреться":
		return player.lookAround()
	case "идти":
		return player.move(parts[1])
	case "надеть":
		/*player.inventory = append(player.inventory, "рюкзак")
		rooms["комната"].description[0] = "на столе: ключи, конспекты. можно пройти - коридор"
		player.currentRoom.changeItems("рюкзак")*/
		return player.putOn(parts[1])
	case "взять":
		return player.take(parts[1])
	case "применить":
		return player.use(parts[1], parts[2])
	default:
		return "неизвестная команда"
	}
}

func main() {
	initGame()

}
