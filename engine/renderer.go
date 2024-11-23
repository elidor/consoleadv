package engine

// A renderer "draws" on a canvas
type Renderer interface {
	Render(*Canvas)
}

type Canvas interface {
}

type TextCanvas struct {
}

type TextRenderer struct {
}

func (textRenderer *TextRenderer) Render(canvas *Canvas) {

}
