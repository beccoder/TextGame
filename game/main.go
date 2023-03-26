package main

import (
	"strings"
	"fmt"
)
/*
	код писать в этом файле
	наверняка у вас будут какие-то структуры с методами, глобальные перменные ( тут можно ), функции
*/


type Room struct {
	Name string
	InitMessage string
	NextRoom []string
	InRoom map[string][]string
}

type Person struct {
	// CurrentRoom string
	Bag bool
	Have []string
	Mission []string
	// Rooms map[string]Room
}

func (person *Person) Glance() string {
		var result string
		result += room.InitMessage
		if len(room.InRoom) > 0 {
			for place, items := range room.InRoom {
				result = result + ", " + place + ": "
				for index, item := range items {
					if index == 0 {
						result += item
					} else {
						result += ", " + item
					}
				}
			}
		}
		if room.Name == "kitchen" {
			result += ", " + person.Mission
			// result += ", надо "
			// for index, mission := range person.Mission {
			// 	if index == 0 {
			// 		result += mission
			// 	} else {
			// 		result += " и " + mission
			// 	}
			// } 
		}
		if len(room.NextRoom) > 0 {
			result += ". можно пройти - "
			for index, path := range room.NextRoom {
				if index == 0 {
					result += path
				} else {
					result += ", " + path
				}
			}
		}
}

func (room *Room) Move() {
	// 
}

func (room *Room) Wear() {
	
}

func (room *Room) Take() {
	// 
}

func (room *Room) Apply {
	// 
}

func main() {
	/*
		в этой функции можно ничего не писать
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все команты
		если что-то было - оно корректно перезатирается
	*/
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	var result string
	
	return "not implemented"
}
