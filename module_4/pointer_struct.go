package main

import "fmt"

type Home struct {
	IsOpenDoor bool
}

// OpenDoor - открывает дверь
func (h Home) OpenDoor() {
	h.IsOpenDoor = true
}

func main() {

	home := Home{
		IsOpenDoor: false,
	}

	// По умолчанию дверь закрыта,
	// убедимся, что это так
	fmt.Println(home)

	// Вызывем метод OpenDoor().
	// Отрываем дверь
	home.OpenDoor()

	// Смотрим на результат.
	fmt.Println(home)
}
