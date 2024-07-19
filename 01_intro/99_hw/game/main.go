package main

import (
	"fmt"
	"strings"
)

//yakwlik

type Room struct { //описываем команту
	name        string            //хранит название команты
	description string            //хранит описание команты
	exits       map[string]string //карта в какую команту можно выйти, где ключ - название комнаты, значение - описание команты
}

type Player struct { //описываем игрока
	currentRoom *Room //описывает комнату, в которой находится сейчас игрок
	//inventory   map[string]bool //описывает инвентарь игрока
	inventory []string
}

var rooms map[string]*Room //создаем мапу, которая хранит все команты нашей игры
var player Player          //создаем игрока
// нужно описать дейтсвия - осмотреться(+), идти(+),надеть(-),взять(-),применить(-)

func (p *Player) lookAround() string { //создали метод  - осмотреться
	return p.currentRoom.description //возвращает описание комнаты в которой сейчас находится игрок
}

func (p *Player) move(direction string) string { // создали метод идти
	nextRoomName := p.currentRoom.exits[direction]
	if nextRoomName == "" {
		return "Из этой комнты сюда попасть нельзя"
	}
	p.currentRoom = rooms[nextRoomName]
	return p.lookAround()
}

func (p *Player) take(item string) string {
	p.inventory = append(p.inventory, item)
	return "предмет добавлен в инвентарь: " + item
}
func initGame() {
	rooms = make(map[string]*Room) //инициализируем карту

	// далее создаем каждую команту
	rooms["кухня"] = &Room{
		name:        "Кухня",
		description: "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
		exits:       map[string]string{"коридор": "коридор"},
	}

	rooms["коридор"] = &Room{
		name:        "Коридор",
		description: "ничего интересного. можно пройти - кухня, комната, улица",
		exits:       map[string]string{"кухня": "кухня", "комната": "комната", "улица": "улица"},
	}

	rooms["комната"] = &Room{
		name:        "Комната",
		description: "ты в своей комнате. можно пройти - коридор",
		exits:       map[string]string{"коридор": "коридор"},
	}
	rooms["улица"] = &Room{
		name:        "Улица",
		description: "на улице весна. можно пройти - домой",
		exits:       map[string]string{"кухня": "кухня"},
	}

	player = Player{
		currentRoom: rooms["кухня"],
		inventory:   make([]string, 5, 10)} //создаем игрока и устанавливаем ему начальное положение
}

func handleCommand(command string) string {
	command = strings.ToLower(command) //приводим команду к нижнему регистру
	//это тестовые свитчи
	/*
		в идеале command разбивать сплитом, и уже по первому дейсвию command[0] идти в кейсы
		это я доделаю, но сначала нужно разобраться с другими командами
	*/
	switch command {
	case "осмотреться":
		return player.lookAround()
	case "идти коридор", "идти комната", "идти кухня", "идти улица":
		parts := strings.Split(command, " ") //разделяем команду на  дейсвтие-куда
		return player.move(parts[1])         //вызываем метод и идем в эту комнату
	case "надеть рюкзак":
		return "вы надели: рюкзак"
	case "взять ключи", "взять конспекты":
		item := strings.Split(command, " ")
		return player.take(item[1])
	case "применить ключи дверь":
		return "дверь открыта"
	default:
		return "Неизвестная команда."
	}
}

func main() {
	initGame()
	/*fmt.Println(handleCommand("завтракать"))
	fmt.Println("начальное положение ", player.currentRoom.name)
	fmt.Println("получаем команду - осмотреться")
	fmt.Println(player.lookAround())
	fmt.Println("команда выполнилась")
	fmt.Println("получаем команду - идти коридор")
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println("команда выполнилась")
	fmt.Println("теперь мы в комнате ", player.currentRoom.name)
	fmt.Println("=====")
	fmt.Println("получаем команду - идти комната")
	fmt.Println(handleCommand("идти комната"))
	fmt.Println("команда выполнилась")
	fmt.Println("теперь мы в комнате ", player.currentRoom.name)
	fmt.Println("=====")
	//fmt.Println(handleCommand("идти коридор"))
	//fmt.Println(handleCommand("идти кухня"))
	fmt.Println(handleCommand("идти улица"))*/
	handleCommand("взять ключи")
	handleCommand("взять конспекты")
	fmt.Println(player.inventory[0])

}
