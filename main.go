package main

import (
	"image/color"
	"log"

	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"

	"github.com/hajimehoshi/ebiten"

	"github.com/peterhellberg/gfx"
)

type Hoverable interface {
	IsAbove(gfx.Vec) bool
	Hovered(bool)
}

const (
	screenWidth  = 600
	screenHeight = 600
)

func cursorPosition() gfx.Vec {
	x, y := ebiten.CursorPosition()
	return gfx.V(float64(x), float64(y))
}

var (
	sampleText      = `The quick brown fox jumps over the lazy dog.`
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

type Menu struct {
	gfx.Rect
	hovered bool
}

func (m *Menu) render(screen *ebiten.Image, i int) {
	var c color.RGBA
	switch i {
	case 0:
		c = colornames.Sienna
	case 1:
		c = colornames.Tomato
	case 2:
		c = colornames.Greenyellow
	case 3:
		c = colornames.Teal
	}

	if m.hovered {
		c = colornames.White
	}

	ebitenutil.DrawRect(screen, m.Min.X, m.Min.Y, m.W(), m.H(), c)
}

var menus []Menu

func update(screen *ebiten.Image) error {

	menu.hover(cursorPosition())
	menu.render(screen)
	return nil
}

var longText = "Your crew boards the station, cautiously moving between corridors. Suddenly a man-sized arachnid bursts from a vent in the ceiling, followed by countless more. You fight your way back to the airlock and are forced to leave before accounting for all crew members. Not everybody made it back."

var menu *Menu2

func main() {
	xMin, yMin := 60.0, 60.0
	base := gfx.R(0, 0, 100, 50)
	for y := 0.0; y < 2; y++ {
		for x := 0.0; x < 2; x++ {
			menus = append(menus, Menu{base.Moved(gfx.V(xMin+10*x, yMin+10*y)), false})
		}
	}

	menu = NewMenu2(
		gfx.R(50, 50, screenWidth-50, screenHeight-50),
		TextBox{text: longText},
		TextBox{text: "1. Fight the monster. This is a very dangerous option, so the text is a bit longer"},
		TextBox{text: "2. Try to escape"},
	)

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "abomination"); err != nil {
		log.Fatal(err)
	}
}

func drawLines(screen *ebiten.Image, lines []string, fnt font.Face, x, y int, c color.Color) {
	offset := fnt.Metrics().Height.Ceil()

	for _, line := range lines {
		text.Draw(screen, line, fnt, x, y+offset, c)
		offset += fnt.Metrics().Height.Ceil()
	}
}

func drawRect(screen *ebiten.Image, r gfx.Rect, thickness float64, c color.Color) {
	// Extend lines to make proper corners
	gfx.DrawLine(screen, r.Min, r.Min.AddXY(r.W()+thickness, 0), thickness, c)
	gfx.DrawLine(screen, r.Max.AddXY(0, thickness), r.Max.AddXY(0, -r.H()), thickness, c)
	gfx.DrawLine(screen, r.Max, r.Max.AddXY(-r.W()-thickness, 0), thickness, c)
	gfx.DrawLine(screen, r.Min.AddXY(0, -thickness), r.Min.AddXY(0, r.H()), thickness, c)
}
