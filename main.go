package main

import (
	"github.com/magarcia/intel8080/io"
)

func emulate(bytes []byte) {

}

func main() {
	rom, err := io.LoadROM("invaders/invaders.h")

	if err != nil {
		panic(err)
	}
	emulate(rom)
}
