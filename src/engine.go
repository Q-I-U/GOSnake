package game

import(
	"github.com/veandco/go-sdl2/sdl"
	sdlttf "github.com/veandco/go-sdl2/ttf"
	"time"
)

type Rectangle struct {
	x, y, x2, y2, x3, y3 int32
}

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
	e.MenuInput()
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

func board(e *engine, r *sdl.Renderer){
	var rect sdl.Rect
	r.Clear()
	r.SetDrawColor(0, 0, 0, 255)
	rect = sdl.Rect{0, 0, 800, 600}
	r.FillRect(&rect)
	r.SetDrawColor(255, 255, 255, 255)
	for i := 50; i < e.c.width; i = i + 50{
		r.DrawLine(int32(i), 0, int32(i), 600)
	}
	for j := 50; j < e.c.height; j = j + 50{
		r.DrawLine(0, int32(j), 800, int32(j))
	}
}

func New(r *sdl.Renderer) Rectangle{
	pos := Rectangle{100, 50, 50, 50, 0, 50}
	rect := []sdl.Rect{{pos.x, pos.y, 50, 50}, {pos.x2, pos.y2, 50, 50}, {pos.x3, pos.y3, 50, 50}}
	r.SetDrawColor(0, 173, 17, 255)
	r.FillRects(rect)
	return pos
}

func movePlayer(pos Rectangle, r *sdl.Renderer, e *engine, direction string) Rectangle{
	if direction == "right" && pos.x + 50 < 800{
		pos = Rectangle{pos.x + 50, pos.y, pos.x2 + 50, pos.y2, pos.x3 + 50, pos.y3} 
	}else if direction == "left" && pos.x - 50 >= 0{
		pos = Rectangle{pos.x - 50, pos.y, pos.x2 - 50, pos.y2, pos.x3 - 50, pos.y3} 
	}else if direction == "up" && pos.y - 50 >= 0{
		pos = Rectangle{pos.x, pos.y - 50, pos.x2, pos.y2 - 50, pos.x3, pos.y3 - 50} 
	}else if direction == "down" && pos.y + 50 < 600{
		pos = Rectangle{pos.x, pos.y + 50, pos.x2, pos.y2 + 50, pos.x3, pos.y3 + 50} 
	}
	board(e, r)
	rect := []sdl.Rect{{pos.x, pos.y, 50, 50}, {pos.x2, pos.y2, 50, 50}, {pos.x3, pos.y3, 50, 50}}
	r.SetDrawColor(0, 173, 17, 255)
	r.FillRects(rect)
	r.Present()
	return pos
}

func (e *engine) MenuInput(){
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
					Start(e)
				}
			}
		}
	}
}

func (e *engine) GameInput(pos Rectangle){
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
					case sdl.K_RIGHT:
						pos = movePlayer(pos, e.render, e, "right")
						time.Sleep(1 * time.Second)
					case sdl.K_LEFT:
						pos = movePlayer(pos, e.render, e, "left")
						time.Sleep(1 * time.Second)
					case sdl.K_UP:
						pos = movePlayer(pos, e.render, e, "up")
						time.Sleep(1 * time.Second)
					case sdl.K_DOWN:
						pos = movePlayer(pos, e.render, e, "down")
						time.Sleep(1 * time.Second)
					}
				}
			}
		}
}