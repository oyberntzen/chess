package main

import (
	"log"

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

func main() {
	ebiten.SetWindowSize(400, 400)
	g := game{}
	g.board = bitboard.FenString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -")

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
