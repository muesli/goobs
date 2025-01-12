package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"
)

func main() {
	client, err := goobs.New(os.Getenv("WSL_HOST")+":4444", goobs.WithPassword("hello"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()

	version, _ := client.General.GetVersion()
	fmt.Printf("Websocket server version: %s\n", version.ObsWebsocketVersion)
	fmt.Printf("OBS Studio version: %s\n", version.ObsStudioVersion)

	// This event loop is in the foreground now. If you mess around in OBS,
	// you'll see different events popping up!

	for event := range client.IncomingEvents {
		switch e := event.(type) {
		case *events.Error:
			log.Printf("Got an error: %s", e.Err)
		default:
			log.Printf("Unhandled event: %#v", e.GetUpdateType())
		}
	}
}
