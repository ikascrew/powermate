package main

import (
	"fmt"
	"log"

	pm "github.com/ikascrew/powermate"
)

func main() {
	pm.HandleFunc(trigger)
	go func() {
		err := pm.Listen("/dev/input/powermate")
		if err != nil {
			log.Fatal(err)
		}
	}()
	var stdin string
	fmt.Scan(&stdin)
}

func trigger(e pm.Event) error {
	switch e.Type {
	case pm.Rotation:
		switch e.Value {
		case pm.Left:
			fmt.Println("Left")
		case pm.Right:
			fmt.Println("Right")
		}
	case pm.Press:
		switch e.Value {
		case pm.Up:
			fmt.Println("Up")
		case pm.Down:
			fmt.Println("Down")
		}
	default:
		fmt.Println("Default")
	}
	return nil
}
