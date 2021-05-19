package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten"
	"github.com/oyberntzen/chessbot/bitboard"
	"github.com/oyberntzen/chessbot/graphics"
)

type game struct {
	board bitboard.ChessBoard
}

//Update handles the logic
func (g *game) Update(screen *ebiten.Image) error {
	graphics.HandleInput(&g.board)
	return nil
}

//Draw handles displaying each frame
func (g *game) Draw(screen *ebiten.Image) {
	graphics.DrawBoard(screen)
	graphics.DrawPieces(screen, &g.board)
}

//Layout returns the size of the canvas
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	ebiten.SetWindowSize(400, 400)
	g := game{}
	g.board = bitboard.FenString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -")
	g.board.Init()
	g.board.InitVariables()

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
