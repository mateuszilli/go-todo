package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type todo struct {
	description string
}

func render(l *[]todo, selected *int) {
	if len(*l) > 0 {
		for index, item := range *l {
			if index == *selected {
				fmt.Print("\u2714")
				fmt.Print("   ")
			} else {
				fmt.Print("    ")
			}

			fmt.Print(item.description)

			fmt.Println()
		}
	} else {
		fmt.Println("Nothing to show...")
	}

	fmt.Println()
}

func main() {
	ch := make(chan byte)
	go func(ch chan byte) {
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

		for {
			buffer := bufio.NewReaderSize(os.Stdin, 1)
			input, _ := buffer.ReadByte()
			ch <- input
		}
	}(ch)

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
		select {
		case stdin, _ := <-ch:
			switch stdin {
			case 10:
				list = append(list[:selected], list[selected+1:]...)
				selected = 0
			case 27:
				// fmt.Println("ESC pressed")
			case 106:
				if selected < len(list)-1 {
					selected++
				}
			case 107:
				if selected > 0 {
					selected--
				}
			case 110:
				for {
					exec.Command("stty", "-F", "/dev/tty", "echo").Run()
					reader := bufio.NewReader(os.Stdin)
					fmt.Print("Write the description of the new task:")
					text, _ := reader.ReadString('\n')
					list = append(list, todo{description: text})
					break
				}
			}

			render(&list, &selected)

		default:
			render(&list, &selected)
		}

		fmt.Println("Press N to add new task")
		fmt.Println("Press J or K to move the marker down or up")
		fmt.Println("Press ENTER to complete the marked task")
		fmt.Println("Press ESC or Ctrl-C to exit this program")
		time.Sleep(time.Millisecond * 100)
		fmt.Print("\033[H\033[2J")
	}
}
