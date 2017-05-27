package powermate

import (
	"encoding/binary"
	"os"
)

type Event struct {
	Type  Type
	Value Value
}

type Type int

const (
	Rotation = 1
	None     = 0
	Press    = -1
)

type Value int

const (
	Left  = 0
	Right = -1
)

const (
	Up   = 0
	Down = -1
)

var handler func(Event) error

func HandleFunc(fn func(Event) error) {
	handler = fn
}

func Listen(fname string) error {

	device, err := os.OpenFile(fname, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer device.Close()

	buf := make([]byte, 48)
	for {
		n, err := device.Read(buf)
		if err != nil {
			return err
		}
		event := buf[:n]

		w := event[16:20]
		typ, _ := binary.Varint(w)
		w = event[20:24]
		val, _ := binary.Varint(w)

		err = handler(Event{
			Type:  Type(typ),
			Value: Value(val),
		})

		if err != nil {
			return err
		}
	}
}
