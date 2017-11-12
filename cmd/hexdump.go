package main

import (
	"fmt"
	"github.com/magarcia/intel8080/io"
)

func hexdump(rom []byte) {

	for i, element := range rom {
		if (i % 16) == 0 {
			fmt.Printf("\n%08x ", i)
		}
		fmt.Printf("%02x", element)
		if (i % 16) != 15 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}

func main() {
	rom, err := io.LoadROM("invaders/invaders.h")

	if err != nil {
		panic(err)
	}
	hexdump(rom)
}
