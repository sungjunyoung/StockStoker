package fzf

import "fmt"

const (
	// https://www2.ccs.neu.edu/research/gpc/MSim/vona/terminal/vtansi.htm
	eraseScreen = "\x1b[2J"
)

type renderer struct {
}

func (r *renderer) clear() {
	r.forceCursorPosition(0, 0)
	fmt.Printf(eraseScreen)
}

func (r *renderer) forceCursorPosition(row int, col int) {
	fmt.Printf("\x1b[%d;%df", row, col)
}

func (r *renderer) renderHeader(header string) {
	fmt.Print(header)
	fmt.Println()
}

func (r *renderer) renderList(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}
