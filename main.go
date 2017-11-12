package main

import (
	"github.com/magarcia/emulator/cmd"
	"github.com/magarcia/emulator/io"
)

func main() {
	rom, err := io.LoadROM("invaders/invaders.h")

	if err != nil {
		panic(err)
	}
	cmd.Disassembler(rom, false, true)
}
