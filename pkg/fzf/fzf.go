package fzf

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"strings"
)

const (
	searchPrefix = "🔎 > "
	redColor     = "\033[1;31m%s\033[0m"
)

type Item interface {
	Id() string
	String() string
}

type Items interface {
	EachItem(handler func(Item))
	ListString() []string
}

type Fzf struct {
	renderer       *renderer
	items          Items
	filteredItems  []Item
	searchingText  string
	searchingIndex int
}

func New(items Items) *Fzf {
	var filteredItems []Item
	if items != nil {
		items.EachItem(func(item Item) {
			filteredItems = append(filteredItems, item)
		})
	}

	return &Fzf{
		renderer:      &renderer{},
		items:         items,
		filteredItems: filteredItems,
	}
}

func (f *Fzf) SelectContent() (Item, error) {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		return nil, err
	}
	defer keyboard.Close()

	f.rerender()
	for {
		event := <-keysEvents
		if event.Err != nil {
			return nil, event.Err
		}

		if f.isQuitInput(event.Key) {
			return nil, fmt.Errorf("no item selected")
		}

		if f.isIgnoreInput(event.Key) {
			continue
		}

		switch event.Key {
		case keyboard.KeyBackspace, keyboard.KeyBackspace2:
			f.deleteSearchingText()
		case keyboard.KeyArrowUp:
			f.decreaseIndexInFilteredItems()
		case keyboard.KeyArrowDown:
			f.increaseIndexInFilteredItems()
		case keyboard.KeyEnter:
			f.renderer.clear()
			return f.filteredItems[f.searchingIndex], nil
		default:
			f.searchingText += string(event.Rune)
		}

		f.rerender()
	}
}

func (f *Fzf) rerender() {
	f.renderer.clear()
	f.renderer.renderHeader(searchPrefix + f.searchingText)

	f.setFilteredItems()
	var items []string
	for i, item := range f.filteredItems {
		if i == f.searchingIndex {
			items = append(items, fmt.Sprintf(fmt.Sprintf(redColor, "> ")+"%s", item.String()))
			continue
		}
		items = append(items, item.String())
	}
	f.renderer.renderList(items)

	f.renderer.forceCursorPosition(0, len(searchPrefix)+len(f.searchingText))
}

func (f *Fzf) setFilteredItems() {
	var filteredItems []Item

	f.items.EachItem(func(item Item) {
		if strings.Contains(item.String(), f.searchingText) {
			filteredItems = append(filteredItems, item)
		}
	})

	f.filteredItems = filteredItems
}

func (f *Fzf) isQuitInput(key keyboard.Key) bool {
	return key == keyboard.KeyEsc ||
		key == keyboard.KeyCtrlC ||
		key == keyboard.KeyCtrlD
}

func (f *Fzf) isIgnoreInput(key keyboard.Key) bool {
	return key == keyboard.KeyArrowLeft ||
		key == keyboard.KeyArrowRight
}

func (f *Fzf) deleteSearchingText() {
	if len(f.searchingText) < 1 {
		f.searchingText = ""
		return
	}
	r := []rune(f.searchingText)
	f.searchingText = string(r[:len(r)-1])
}

func (f *Fzf) decreaseIndexInFilteredItems() {
	if f.searchingIndex-1 < 0 {
		f.searchingIndex = 0
		return
	}
	f.searchingIndex -= 1
}

func (f *Fzf) increaseIndexInFilteredItems() {
	if f.searchingIndex+1 >= len(f.filteredItems) {
		f.searchingIndex = len(f.filteredItems) - 1
		return
	}
	f.searchingIndex += 1
}
