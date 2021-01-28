package main

import(
	"github.com/GNUSheep/GOsnake/src"
)

func main(){
	conf := Config{800, 600}
	err := game.Run(&conf)
	if err != nil{
		return err
	}
}