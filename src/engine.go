package game

import(
	"github.com/veandco/go-sdl2/sdl"
	sdlttf "github.com/veandco/go-sdl2/ttf"
	"fmt"
)

type Engine interface{
	Start()
}

type engine struct{
	c *Config
	win *sdl.Window
	render *sdl.Renderer
}

type render struct{
	r *sdl.Renderer
	c *Config
}

func NewEngine(c *Config) (Engine, error){
	win, render, err := sdl.CreateWindowAndRenderer(int32(c.width), int32(c.height), sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	win.SetTitle("Snake")

	return &engine{c, win, render}, nil
}

func (e *engine) Start(){
	menu(e.render, e.win)
	e.KeyboardInput()
}

func menu(r *sdl.Renderer, win *sdl.Window){
	var solid *sdl.Surface
	var font *sdlttf.Font
	var surface *sdl.Surface

	font, _ = sdlttf.OpenFont("./assets/dejavu.ttf", 32)
	surface, _ = win.GetSurface()
	solid, _ = font.RenderUTF8Solid("              Press ENTER to play or Q to quit.", sdl.Color{255, 255, 255, 255})
	solid.Blit(nil, surface, nil)
	win.UpdateSurface()
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
				switch event.(*sdl.KeyboardEvent).Keysym.Sym {
				case sdl.K_q:
					break loop
				case sdl.K_RETURN:
					fmt.Println("Cooming soon!")
					break loop
				}
			}
		}
	}
}