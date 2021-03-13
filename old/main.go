package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const king uint8 = 1
const queen uint8 = 2
const bishop uint8 = 3
const knight uint8 = 4
const rook uint8 = 5
const pawn uint8 = 6

const white uint8 = 8
const black uint8 = 16

const cellSize int = 50

var board chessBoard

type game struct{}

//Update handles the logic
func (g *game) Update(screen *ebiten.Image) error {
	board = handleInput(board)
	return nil
}

//Draw handles displaying each frame
func (g *game) Draw(screen *ebiten.Image) {
	drawBoard(screen, board)
	drawPieces(screen, board)
}

//Layout returns the size of the canvas
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	board = fenToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")

	ebiten.SetWindowSize(400, 400)
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
