package engine

import "fmt"

func Run() {
	renderer := TextRenderer{}
	canvas := fmt.Println
	for {
		renderer.Render(canvas)
	}
}
