package main

import(
	"fmt"
	"github.com/GNUSheep/GOsnake/src"
)

func main() {
	conf := game.DefaultConfig
	err := game.Run(&conf)
	if err != nil{
		fmt.Println("Error")
	}
}