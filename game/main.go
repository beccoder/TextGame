package main

import (
	"strings"
	"fmt"
)

type Room struct {
	Name string
	InitMessage string
	Status string
	NextRoom []string
	InRoom map[string][]string
}

func (room Room) getNextRooms() string {
	var result string
	if len(room.NextRoom) > 0 {
		result += "можно пройти - "
		for index, path := range room.NextRoom {
			if index > 0 {
				result += ", "
			}
			result += path
		}
	}
	return result
}

func (room Room) getItemsInRoom() string {
	var result string
	if len(room.InRoom) > 0 {
		var c int
		for place, items := range room.InRoom {
			if c++; c > 1 {
				result += ", "
			}
			result = result + place + ": "
			for index, item := range items {
				if index > 0 {
					result += ", "
				}
				result += item
			}
		}
	} else {
		switch room.Name {
			case "кухня", "комната", "улица":
				result += "пустая " + room.Name
			case "коридор":
				result += "пустой " + room.Name
		}
	}
	return result
}

func (room Room) isInRoom(item string) bool {
	for _, vals := range room.InRoom {
		for _, val := range vals {
			if val == item {
				return true
			}
		}
	}
	return false
}

func (room *Room) delFromRoom(item string) {
	for key, vals := range room.InRoom {
		for index, val := range vals {
			if val == item {
				room.InRoom[key] = append(vals[:index], vals[index+1:]...)
				if len(room.InRoom[key]) == 0 {
					delete(room.InRoom, key)
				}
			}
		}
	}
}

type Player struct {
	CurrentRoom string
	Rooms map[string]Room
	DoorIsOpen bool
	HasBag bool
	Have []string
	Mission string
}

func (plr *Player) Glance() string {
	var result string
	room := plr.Rooms[plr.CurrentRoom]
	result += room.getItemsInRoom()
	if room.InitMessage != "" {
		result = room.InitMessage + ", " + result + ", " + plr.Mission
	}
	result += ". " + room.getNextRooms()
	return result
}

func contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func (plr *Player) Move(roomName string) string {
	var result string
	room := plr.Rooms[plr.CurrentRoom]
	if _, ok := plr.Rooms[roomName]; ok {
		if contains(room.NextRoom, roomName) {
			if roomName == "улица" && plr.DoorIsOpen == false {
				return "дверь закрыта"
			}
			plr.CurrentRoom = roomName
			room := plr.Rooms[plr.CurrentRoom]
			result = room.Status + ". " + room.getNextRooms()
		} else {
			result = "нет пути в " + roomName
		}
	} else {
		result = "нет такая комната"
	}
	return result
}

func (plr *Player) Wear(item string) string {
	var result = "нет такого"
	if item == "рюкзак" && plr.HasBag == false {
		room := plr.Rooms[plr.CurrentRoom]
		if room.isInRoom(item) {
			plr.HasBag = true
			plr.Mission = "надо идти в универ"
			room.delFromRoom(item)
			result = "вы надели: " + item
		}
	}
	return result
}

func (plr *Player) Take(item string) string {
	var result string
	room := plr.Rooms[plr.CurrentRoom]
	if room.isInRoom(item) {
		if plr.HasBag {
			plr.Have = append(plr.Have, item)
			room.delFromRoom(item)
			result = "предмет добавлен в инвентарь: " + item
		} else {
			result = "некуда класть"
		}
	} else {
		result = "нет такого"
	}
	return result
}

func (plr *Player) Apply(data []string) string {
	var result string
	key := data[0]
	whereToApply := data[1]
	if contains(plr.Have, key) {
		if key == "ключи" && whereToApply == "дверь" {
			plr.DoorIsOpen = true
			result = "дверь открыта"
		} else {
			result = "не к чему применить"
		}
	} else {
		result = "нет предмета в инвентаре - " + key
	}
	return result
}

var player Player

func main() {
	initGame()
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println(handleCommand("идти комната"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("надеть рюкза"))
	fmt.Println(handleCommand("взять ключи"))
	fmt.Println(handleCommand("взять конспекты"))
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println(handleCommand("применить ключи дверь"))
	fmt.Println(handleCommand("идти улица"))
	
	initGame()
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("завтракать"))
	fmt.Println(handleCommand("идти комната"))
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println(handleCommand("применить ключи дверь"))
	fmt.Println(handleCommand("идти комната"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("взять ключи"))
	fmt.Println(handleCommand("надеть рюкзак"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("взять ключи"))
	fmt.Println(handleCommand("взять телефон"))
	fmt.Println(handleCommand("взять ключи"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("взять конспекты"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println(handleCommand("идти кухня"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println(handleCommand("идти улица"))
	fmt.Println(handleCommand("применить ключи дверь"))
	fmt.Println(handleCommand("применить телефон шкаф"))
	fmt.Println(handleCommand("применить ключи шкаф"))
	fmt.Println(handleCommand("идти улица"))
}

func initGame() {
	fmt.Printf(`
	*****Welcome to the Text Game!*****

	There is 5 kind of actions in our game:
	---Glance(осмотреться) - to look through rooms---
	---Move(идти) - to move another room---
	---Wear(надеть) - to wear an item---
	---Take(взять) - to load an item into a bag---
	---Apply(применить) - to use an item---
	`)

	player = Player{
		CurrentRoom: "кухня",
		Rooms: map[string]Room{
			"кухня": Room{
				Name: "кухня",
				InitMessage: "ты находишься на кухне",
				Status: "кухня, ничего интересного",
				NextRoom: []string{"коридор"},
				InRoom: map[string][]string{
					"на столе": []string{"чай"},
				},
			},
			"коридор": Room{
				Name: "коридор",
				InitMessage: "",
				Status: "ничего интересного",
				NextRoom: []string{"кухня", "комната", "улица"},
				InRoom: map[string][]string{},
			},
			"комната": Room{
				Name: "комната",
				InitMessage: "",
				Status: "ты в своей комнате",
				NextRoom: []string{"коридор"},
				InRoom: map[string][]string{
					"на столе": []string{"ключи", "конспекты"},
					"на стуле": []string{"рюкзак"},
				},
			},
			"улица": Room{
				Name: "улица",
				InitMessage: "",
				Status: "на улице весна",
				NextRoom: []string{"домой"},
				InRoom: map[string][]string{},
			},
		},
		DoorIsOpen: false,
		HasBag: false,
		Have: []string{},
		Mission: "надо собрать рюкзак и идти в универ",
	}
}

func handleCommand(command string) string {
	buffer := strings.Split(command, " ")
	l := len(buffer)
	switch {
		case buffer[0] == "осмотреться" && l == 1:
			return player.Glance()
		case buffer[0] == "идти" && l == 2:
			return player.Move(buffer[1])
		case buffer[0] == "надеть" && l == 2:
			return player.Wear(buffer[1])
		case buffer[0] == "взять" && l == 2:
			return player.Take(buffer[1])
		case buffer[0] == "применить" && l == 3:
			return player.Apply([]string{buffer[1], buffer[2]})
		default:
			return "неизвестная команда"
	}
}