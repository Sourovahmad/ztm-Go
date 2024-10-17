package main

import "fmt"

type Room struct {
	Name      string
	IsCleaned bool
}

func checkIfRoomIsCleaned(room Room) {

	if room.IsCleaned {
		fmt.Println("the room has been cleaned")
	} else {
		fmt.Println("the room has not cleaned")
	}
}

func main() {

	rooms := [...]Room{
		{
			Name:      "office",
			IsCleaned: false,
		},
		{
			Name:      "dressin room",
			IsCleaned: true,
		},

		{
			Name:      "shower room",
			IsCleaned: false,
		},
	}

	for _, room := range rooms {
		checkIfRoomIsCleaned(room)
	}

	fmt.Println("hello this message is from vim")

}
