package keyboardUtils

import (
	"computer/src/bus"
	"fmt"

	"github.com/eiannone/keyboard"
)

func KeyboardInit() {
	if err := keyboard.Open(); err != nil {
		fmt.Printf("Failed to open keyboard: %s\n", err)
		return
	}
}

func HandleKeyboard(bus *bus.Bus) (bool, bool) {

	var step = false
	var running = true

	char, _, err := keyboard.GetKey()
	if err != nil {
		fmt.Printf("Error reading key: %s\n", err)
		return false, false
	}

	switch char {
	case 'd':
		fmt.Println("D key pressed. Debugging...")
		fmt.Printf("0x%02X", bus.Read(0x6000))
	case 's':
		step = true
	case 'q':
		fmt.Println("Q key pressed. Exiting...")
		running = false
	}
	return running, step
}

func Close() {
	keyboard.Close()
}
