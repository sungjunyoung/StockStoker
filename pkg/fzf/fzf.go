package fzf

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"strings"
)

const (
	searchPrefix = "🔎 > "

	// https://www2.ccs.neu.edu/research/gpc/MSim/vona/terminal/vtansi.htm
	eraseScreen = "\x1b[2J"
)

type Fzf struct {
	contentsToSearch []string
}

func New(contentsToSearch []string) *Fzf {
	return &Fzf{contentsToSearch: contentsToSearch}
}

func (f *Fzf) rerender(searchText string) {
	var contents []string
	for _, content := range f.contentsToSearch {
		if strings.Contains(content, searchText) {
			contents = append(contents, content)
		}
	}

	fmt.Print(f.forceCursorPosition(0, 0))
	fmt.Print(eraseScreen)
	fmt.Print(searchPrefix)
	fmt.Print(searchText)
	fmt.Println()
	for _, content := range contents {
		fmt.Println(content)
	}
	fmt.Print(f.forceCursorPosition(0, len(searchPrefix)+len(searchText)))
}
func (f *Fzf) forceCursorPosition(row int, col int) string {
	return fmt.Sprintf("\x1b[%d;%df", row, col)
}

func (f *Fzf) Run() {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	f.rerender("")

	searchText := ""
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		if event.Key == keyboard.KeyEsc ||
			event.Key == keyboard.KeyCtrlC ||
			event.Key == keyboard.KeyCtrlD {
			break
		}

		switch event.Key {
		case keyboard.KeyBackspace, keyboard.KeyBackspace2:
			if len(searchText) < 1 {
				searchText = ""
			} else {
				searchText = searchText[:len(searchText)-1]
			}
		default:
			searchText += string(event.Rune)
		}

		f.rerender(searchText)
		//fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
		//fmt.Println(searchText)

	}

}
