package game

import(
	"github.com/veandco/go-sdl2/sdl"
)

type Config struct{
	width int
	height int
}

func Run(c *Config) error{
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()
	engine, err := NewEngine(c)
	if err != nil {
		return err
	}
	engine.Run()
	return nil
}