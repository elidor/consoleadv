package engine

func Run() {
	renderer := TextRenderer{}
	for {
		renderer.Render(TextCanvas{})
	}
}
