package main

import (
	"fmt"
	"github.com/sungjunyoung/StockStoker/pkg/fzf"
	"log"
)

type Item struct {
	title string
	code  string
}

func (i Item) Id() string {
	return i.title
}

func (i Item) String() string {
	return fmt.Sprintf("%s    \t%s", i.title, i.code)
}

type Items []Item

func (is Items) ListString() []string {
	var strs []string
	for _, i := range is {
		strs = append(strs, i.String())
	}
	return strs
}

func (is Items) EachItem(handler func(fzf.Item)) {
	for _, i := range is {
		handler(i)
	}
}

func main() {
	items := Items{
		{
			title: "네이버",
			code:  "123123123",
		},
		{
			title: "카카오",
			code:  "55123123",
		},
		{
			title: "삼성전자",
			code:  "64423",
		},
		{
			title: "Nanox",
			code:  "123123123",
		},
	}
	f := fzf.New(items)
	selected, err := f.SelectContent()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(selected.Id())

	//cmd.Execute()
}
