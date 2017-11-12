package main

import (
	"github.com/magarcia/intel8080/cmd"
	"github.com/magarcia/intel8080/io"
)

func main() {
	rom, err := io.LoadROM("invaders/invaders.h")

	if err != nil {
		panic(err)
	}
	cmd.Disassembler(rom, false, true)
}
