package day14restroomredoubt

import (
	"bytes"
	"fmt"
	. "image"
	"image/color"
	"strings"
	"time"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	Day string = "14"
)

var (
	MaxX int = 11
	MaxY int = 7

	mplusFaceSource *text.GoTextFaceSource
)

type Quadrant int

type Game struct {
	bots     []*Bot
	ctrMoves int
}

type Bot struct {
	Location Point
	Vector   Point
}

func (b *Bot) Move(nrSteps int) {
	newLoc := b.Location.Add(b.Vector.Mul(nrSteps))

	b.Location = Pt(
		((newLoc.X%MaxX)+MaxX)%MaxX,
		((newLoc.Y%MaxY)+MaxY)%MaxY,
	)
}

func CalculateSafetyFactorAfterMoves(bots []Bot, nrMoves int) int {
	var quadrants = make(map[Quadrant]int, 5)

	for _, bot := range bots {
		bot.Move(nrMoves)
		fmt.Printf("bot.Location.String(): %v\n", bot.Location.String())
		quadrants[LocationToQuadrant(bot.Location)] += 1
	}

	var factor = 1
	for idx := 0; idx < 4; idx++ {
		fmt.Printf("quadrants[Quadrant(%d)]: %v\n", idx, quadrants[Quadrant(idx)])
		factor *= quadrants[Quadrant(idx)]
	}

	return factor
}

func LocationToQuadrant(p Point) Quadrant {
	middleX := (MaxX - 1) / 2
	middleY := (MaxY - 1) / 2

	if p.X == middleX || p.Y == middleY {
		return -1
	}
	q := Quadrant(0)

	if p.X > middleX {
		q++
	}
	if p.Y > middleY {
		q += 2
	}

	return q
}

func ParseInput(lines []string) []Bot {
	var bots = make([]Bot, 0, len(lines))

	for _, line := range lines {
		bots = append(bots, parseLine(line))
	}

	return bots
}

func parseLine(line string) Bot {
	parts := strings.Split(line, " ")

	bot := Bot{
		Location: partToPoint(parts[0]),
		Vector:   partToPoint(parts[1]),
	}

	return bot
}

func partToPoint(part string) Point {
	parts := strings.Split(part, "=")
	pointParts := strings.Split(parts[1], ",")

	return Pt(
		utils.ConvStrToI(pointParts[0]),
		utils.ConvStrToI(pointParts[1]),
	)
}

func (g *Game) Update() error {
	g.ctrMoves++
	for _, bot := range g.bots {
		bot.Move(1)
	}

	if g.ctrMoves > 6575 {
		if g.ctrMoves == 6578 {
			time.Sleep(10 * time.Second)
		}
		time.Sleep(1 * time.Second)
	} else {
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	green := color.RGBA{
		R: 10,
		G: 255,
		B: 50,
		A: 255,
	}

	const (
		normalFontSize = 12
		bigFontSize    = 48

		x = 20
	)

	msg := fmt.Sprintf("%d", g.ctrMoves)
	op := &text.DrawOptions{}
	op.GeoM.Translate(100, 20)
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   normalFontSize,
	}, op)

	for _, bot := range g.bots {
		screen.Set(bot.Location.X, bot.Location.Y, green)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MaxX + 30, MaxY
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	bots := ParseInput(lines)

	MaxX = 101
	MaxY = 103

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CalculateSafetyFactorAfterMoves(bots, 100))
}

func solvePart2(inputFile string) {
	MaxX = 101
	MaxY = 103

	lines := utils.ReadFileAsLines(inputFile)

	bots := ParseInput(lines)
	game := &Game{
		bots:     make([]*Bot, 0, len(bots)),
		ctrMoves: 6500,
	}

	for _, b := range bots {
		b.Move(6500)
		game.bots = append(game.bots, &b)
	}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Advent of Code - Restroom Redoubt")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

	var count int = 0
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, count)
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}
	mplusFaceSource = s
}
