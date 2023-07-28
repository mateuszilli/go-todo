package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Todo struct {
	Description string `json:"description"`
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func toJSON(t *[]Todo) {
	bytes, err_parse := json.Marshal(t)
	if err_parse != nil {
		log.Fatal("Error on parse JSON", err_parse)
	}

	err_file := ioutil.WriteFile("todo.json", bytes, 0644)
	if err_file != nil {
		log.Fatal("Error on write JSON file", err_file)
	}
}

func fromJSON() []Todo {
	bytes, err_file := ioutil.ReadFile("./todo.json")
	if err_file != nil {
		log.Fatal("Error on read JSON file", err_file)
	}

	var list []Todo
	err_parse := json.Unmarshal(bytes, &list)
	if err_parse != nil {
		log.Fatal("Error on parse JSON", err_parse)
	}

	return list
}

func main() {
	page := 1
	limit := 10
	selected := 0
	option := byte(0)
	list := fromJSON()

	for {
		clear()

		switch option {
		case 10:
			list = append(list[:selected], list[selected+1:]...)
			selected = 0
			toJSON(&list)
		case 27:
			os.Exit(0)
		case 106:
			if selected < (page*limit)-1 && selected < len(list)-1 {
				selected++
			} else if (page * limit) < len(list) {
				page++
				selected = (page - 1) * limit
			}
		case 107:
			if selected > (page-1)*limit {
				selected--
			} else if page > 1 {
				page--
				selected = (page - 1) * limit
			}
		case 110:
			// enable input buffering and restore the echoing state when exiting
			exec.Command("stty", "-F", "/dev/tty", "sane").Run()
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()

			fmt.Print("Task description: ")
			buffer := bufio.NewReader(os.Stdin)
			read, _ := buffer.ReadString('\n')
			text := strings.TrimSuffix(read, "\n")

			if text != "" {
				list = append(list, Todo{Description: text})
				toJSON(&list)
			}

			clear()
		}

		if len(list) > 0 {
			if page > 1 {
				fmt.Print("\u2b06\ufe0f")
				fmt.Print("   ")
				fmt.Println("Back to page", page-1)
				// fmt.Print("    ")
				fmt.Println()
			}
			for i := (page - 1) * limit; i < page*limit && i < len(list); i++ {
				item := list[i]

				if i == selected {
					fmt.Print("\u2714\ufe0f")
					fmt.Print("   ")
				} else {
					fmt.Print("    ")
				}

				fmt.Println(item.Description)

			}
			if page*limit < len(list) {
				// fmt.Print("    ")
				fmt.Println()
				fmt.Print("\u2b07\ufe0f")
				fmt.Print("   ")
				fmt.Println("Go to page", page+1)

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
