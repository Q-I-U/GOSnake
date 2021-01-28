package game

import(
	"github.com/veandco/go-sdl2"
)

type Engine interface{
	Start()
}

type engine struct{
	c *Config
	win *sdl.Window
	render *sdl.Renderer
}

func NewEngine(c *Config) (engine, error){
	win, render, err := sdl.CreateWindowAndRenderer(int32(c.width), int32(c.height), sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	win.SetTitle("Snake")

	return &engine{c, win, render}, nil
}

func (e *engine) Start(){
	e.KeyboardInput()
}

func (e *engine) KeyboardInput(){
loop:
	for{
		event := sdl.WaitEvent()
		if event == nil{
			sdl.Delay(1000)
			continue
		}
		switch t := event.(type){
		case *sdl.QuitEvent:
			break loop
		case *sdl.KeyboardEvent:
			if t.State == sdl.PRESSED {
				switch ev.(*sdl.KeyboardEvent).Keysym.Sym {
				case sdl.K_q:
					break loop
		}
	}
}

