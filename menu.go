package main

import (
	"image/color"
	"strings"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/textgeometry"
	wordwrap "github.com/mitchellh/go-wordwrap"
	"github.com/peterhellberg/gfx"
)

const borderWidth = 3
const padding = 0

type Menu2 struct {
	gfx.Rect
	parts []*TextBox
}

type TextBox struct {
	gfx.Rect
	text    string
	lines   []string
	padding int
	color   color.Color
}

func NewTextBox(text string, maxWidth, maxHeight float64) *TextBox {
	tbPadding := 15
	tb := TextBox{
		text:    text,
		color:   color.White,
		padding: tbPadding,
	}

	limit := textgeometry.MaxWrapPosition(text, mplusNormalFont, int(maxWidth)-2*tbPadding)
	wrapped := wordwrap.WrapString(text, uint(limit))
	lines := strings.Split(wrapped, "\n")
	_, h := textgeometry.BoundingBox(lines, mplusNormalFont)

	// tb.Rect = gfx.R(0, 0, float64(w+2*tbPadding), float64(h+2*tbPadding))
	tb.Rect = gfx.R(0, 0, maxWidth, float64(h+2*tbPadding))
	tb.lines = lines
	return &tb
}

func (tb *TextBox) render(screen *ebiten.Image, offset gfx.Vec) {
	// drawRect(screen, tb.Rect.Moved(offset), 1.2, colornames.Red)

	pos := tb.Rect.Min.Add(offset)
	drawLines(screen, tb.lines, mplusNormalFont, int(pos.X)+tb.padding, int(pos.Y)+tb.padding, tb.color)
}

func (m *Menu2) hover(cursor gfx.Vec) {

	offset := m.Min.AddXY(padding, padding)
	for i, p := range m.parts {

		if p.Rect.Moved(offset).Contains(cursor) {
			m.parts[i].color = colornames.Yellow
		} else {
			m.parts[i].color = colornames.White
		}

		offset = offset.AddXY(0, p.H())
	}
}

func NewMenu2(boundingBox gfx.Rect, parts ...TextBox) *Menu2 {
	// Subtract the half the border width
	bounds := boundingBox
	bounds.Min = bounds.Min.AddXY(borderWidth/2, borderWidth/2)
	bounds.Max = bounds.Max.AddXY(-borderWidth/2, -borderWidth/2)

	m := Menu2{
		Rect: bounds,
	}

	bounds = boundingBox
	bounds.Min = bounds.Min.AddXY(borderWidth+padding, borderWidth+padding)
	bounds.Max = bounds.Max.AddXY(-(borderWidth + padding), -(borderWidth + padding))

	for _, p := range parts {
		m.parts = append(m.parts, NewTextBox(p.text, bounds.W(), bounds.H()))
	}

	return &m
}

func (m *Menu2) render(screen *ebiten.Image) {
	// drawRect(screen, m.Rect, borderWidth, color.White)

	offset := m.Min.AddXY(padding+borderWidth, padding+borderWidth)
	for _, p := range m.parts {
		p.render(screen, offset)
		offset = offset.AddXY(0, p.H())
	}
}
