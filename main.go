package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type todo struct {
	description string
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	option := byte(0)
	selected := 0
	list := []todo{
		{
			description: "Read articles in english",
		},
		{
			description: "Practice english conversation in Omegle",
		},
		{
			description: "Listen podcasts or watch videos in english",
		},
	}

	for {
		clear()

		switch option {
		case 10:
			list = append(list[:selected], list[selected+1:]...)
			selected = 0
		case 27:
			os.Exit(0)
		case 106:
			if selected < len(list)-1 {
				selected++
			}
		case 107:
			if selected > 0 {
				selected--
			}
		case 110:
			// enable input buffering and restore the echoing state when exiting
			exec.Command("stty", "-F", "/dev/tty", "sane").Run()
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()

			fmt.Print("Task description: ")
			buffer := bufio.NewReader(os.Stdin)
			text, _ := buffer.ReadString('\n')

			if len(text) > 1 {
				list = append(list, todo{description: text})
			}

			clear()
		}

		if len(list) > 0 {
			for index, item := range list {
				if index == selected {
					fmt.Print("\u2714\ufe0f")
					fmt.Print("   ")
				} else {
					fmt.Print("    ")
				}

				fmt.Println(item.description)
			}
		} else {
			fmt.Println("Nothing to do, go take a \U0001f37a")
		}

		fmt.Println()
		fmt.Println("Press N to add new task")
		fmt.Println("Press J or K to move the marker down or up")
		fmt.Println("Press ENTER to complete the marked task")
		fmt.Println("Press ESC or CTRL-C to exit this program")

		// disable input buffering and do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

		buffer := bufio.NewReaderSize(os.Stdin, 1)
		input, _ := buffer.ReadByte()
		option = input
	}
}
