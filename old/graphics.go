package main

import (
	"image"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var darkColor color.RGBA = color.RGBA{168, 121, 101, 255}
var lightColor color.RGBA = color.RGBA{240, 216, 192, 255}
var moveColor color.RGBA = color.RGBA{150, 255, 150, 255}
var markColor color.RGBA = color.RGBA{255, 255, 150, 255}

var pieceImages [12]*ebiten.Image

var mark bool
var marked int
var hold bool

func init() {
	file, _ := os.Open("./pieces.png")
	img, _, _ := image.Decode(file)
	i := 0
	for y := 0; y < 50*2; y += 50 {
		for x := 0; x < 50*6; x += 50 {
			subImg := img.(interface {
				SubImage(r image.Rectangle) image.Image
			}).SubImage(image.Rect(x, y, x+50, y+50))
			pieceImages[i], _ = ebiten.NewImageFromImage(subImg, ebiten.FilterDefault)
			i++
		}
	}
}

func drawBoard(screen *ebiten.Image, board chessBoard) {
	var moves []int
	if mark {
		moves = legalMoves(board, marked, true)
	}
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			isMove := false
			if mark {
				for _, move := range moves {
					if y*8+x == move {
						isMove = true
					}
				}
			}
			if mark && marked == y*8+x {
				ebitenutil.DrawRect(screen, float64(x*50), float64(y*50), 50, 50, markColor)
			} else if isMove {
				ebitenutil.DrawRect(screen, float64(x*50), float64(y*50), 50, 50, moveColor)
			} else if (x+y)%2 == 0 {
				ebitenutil.DrawRect(screen, float64(x*50), float64(y*50), 50, 50, lightColor)
			} else {
				ebitenutil.DrawRect(screen, float64(x*50), float64(y*50), 50, 50, darkColor)
			}
		}
	}
}

func drawPieces(screen *ebiten.Image, board chessBoard) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			piece := board.board[y*8+x]
			if piece > 0 {
				if piece&white > 0 {
					piece -= white + 1

					geoM := ebiten.GeoM{}
					geoM.Translate(float64(x*50), float64(y*50))
					screen.DrawImage(pieceImages[piece], &ebiten.DrawImageOptions{GeoM: geoM})
				} else if piece&black > 0 {
					piece -= black - 5

					geoM := ebiten.GeoM{}
					geoM.Translate(float64(x*50), float64(y*50))
					screen.DrawImage(pieceImages[piece], &ebiten.DrawImageOptions{GeoM: geoM})
				}
			}
		}
	}
}

func handleInput(board chessBoard) chessBoard {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !hold {
			x, y := ebiten.CursorPosition()
			newMarked := (y/50)*8 + (x / 50)
			if mark && newMarked == marked {
				mark = false
			} else {

				var moved bool
				if mark {
					var index int
					for i, j := range board.whitePieces {
						if j == marked {
							index = i
						}
					}
					board, moved = movePiece(board, index, white, newMarked, true)
				}

				mark = true
				if !moved {
					marked = newMarked
				} else {
					mark = false
					board = randomMove(board)
				}
			}
		}
		hold = true
	} else {
		hold = false
	}
	return board
}
