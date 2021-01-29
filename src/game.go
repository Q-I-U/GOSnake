package game

import(
	"github.com/veandco/go-sdl2/sdl"
	sdlttf "github.com/veandco/go-sdl2/ttf"
)

type Config struct{
	width int
	height int
}

var DefaultConfig = Config{
	width:       800,
	height:      600,
}

func Run(c *Config) error{
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	if err := sdlttf.Init(); err != nil {
		return err
	}
	defer sdlttf.Quit()

	engine, err := NewEngine(c)
	if err != nil {
		return err
	}
	engine.Start()
	return nil
}

func Start(e *engine){
	board(e, e.render)
	pos := New(e.render)
	e.render.Present()
	e.GameInput(pos)
}