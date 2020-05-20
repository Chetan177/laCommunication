package fs

import (
	"fmt"
	"github.com/chetan177/go-eventsocket/eventsocket"
	"log"
)

const (
	fsAdd  = "localhost:8021"
	fsPass = "ClueCon"
)

func MakeCall(to, from string) {
	fsSocket, err := eventsocket.Dial(fsAdd, fsPass)
	if err != nil {
		log.Fatal("Error in FS ", err)
	}
	cmd := fmt.Sprintf("bgapi originate user/%s &bridge(user/%s)", from, to)
	_, err = fsSocket.Send(cmd)
	if err != nil {
		log.Fatal("Error in command ", err)
	}
	fsSocket.Close()
}
