package main

import (
	"github.com/sungjunyoung/StockStoker/pkg/fzf"
)

func main() {
	f := fzf.New([]string{"카카오 1392992", "네이버 1202023", "삼성 123123"})
	f.Run()

	//cmd.Execute()
}
