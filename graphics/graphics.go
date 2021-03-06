package graphics

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/oyberntzen/chessbot/bitboard"
)

var darkColor color.RGBA = color.RGBA{168, 121, 101, 255}
var lightColor color.RGBA = color.RGBA{240, 216, 192, 255}
var moveColor color.RGBA = color.RGBA{150, 255, 150, 255}
var markColor color.RGBA = color.RGBA{255, 255, 150, 255}

var pieceImages [12]*ebiten.Image

var marked bitboard.Bitboard
var moves bitboard.Bitboard

var pressed bool

var pawnPromotion bool
var promotionMoves []bitboard.Move

const (
	whitePawnImage   int = 5
	whiteRookImage   int = 4
	whiteKnightImage int = 3
	whiteBishopImage int = 2
	whiteQueenImage  int = 1
	whiteKingImage   int = 0
	blackPawnImage   int = 11
	blackRookImage   int = 10
	blackKnightImage int = 9
	blackBishopImage int = 8
	blackQueenImage  int = 7
	blackKingImage   int = 6
)

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

func drawBitBoard(screen *ebiten.Image, board bitboard.Bitboard, posColor color.RGBA, negColor color.RGBA) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			square := bitboard.CoordsToBitboard(x, y)
			if square&board > 0 {
				ebitenutil.DrawRect(screen, float64(x*50), float64(y*50), 50, 50, posColor)
			} else {
				ebitenutil.DrawRect(screen, float64(x*50), float64(y*50), 50, 50, negColor)
			}
		}
	}
}

func DrawBoard(screen *ebiten.Image) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if (x+y)%2 == 0 {
				ebitenutil.DrawRect(screen, float64(x*50), float64(y*50), 50, 50, lightColor)
			} else {
				ebitenutil.DrawRect(screen, float64(x*50), float64(y*50), 50, 50, darkColor)
			}
		}
	}
	drawBitBoard(screen, marked, markColor, color.RGBA{0, 0, 0, 0})
	drawBitBoard(screen, moves, moveColor, color.RGBA{0, 0, 0, 0})
}

func DrawPieces(screen *ebiten.Image, board *bitboard.ChessBoard) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			square := bitboard.CoordsToBitboard(x, y)
			var img *ebiten.Image
			if board.WhitePawns&square > 0 {
				img = pieceImages[whitePawnImage]
			} else if board.WhiteRooks&square > 0 {
				img = pieceImages[whiteRookImage]
			} else if board.WhiteKnights&square > 0 {
				img = pieceImages[whiteKnightImage]
			} else if board.WhiteBishops&square > 0 {
				img = pieceImages[whiteBishopImage]
			} else if board.WhiteQueens&square > 0 {
				img = pieceImages[whiteQueenImage]
			} else if board.WhiteKing&square > 0 {
				img = pieceImages[whiteKingImage]
			} else if board.BlackPawns&square > 0 {
				img = pieceImages[blackPawnImage]
			} else if board.BlackRooks&square > 0 {
				img = pieceImages[blackRookImage]
			} else if board.BlackKnights&square > 0 {
				img = pieceImages[blackKnightImage]
			} else if board.BlackBishops&square > 0 {
				img = pieceImages[blackBishopImage]
			} else if board.BlackQueens&square > 0 {
				img = pieceImages[blackQueenImage]
			} else if board.BlackKing&square > 0 {
				img = pieceImages[blackKingImage]
			} else {
				continue
			}

			geoM := ebiten.GeoM{}
			geoM.Translate(float64(x*50), float64(y*50))
			screen.DrawImage(img, &ebiten.DrawImageOptions{GeoM: geoM})
		}
	}
}

func HandleInput(board *bitboard.ChessBoard) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !pressed {
			if !pawnPromotion {
				pressed = true
				x, y := ebiten.CursorPosition()
				if x < 400 && y < 400 {
					newMarked := bitboard.CoordsToBitboard(x/50, y/50)
					var psudomoves []bitboard.Move
					var validMove bool
					for _, m := range board.PsudoLegalMoves(false) {
						if m.From == marked && m.To == newMarked {
							psudomoves = append(psudomoves, m)
							validMove = true
						}
					}
					if validMove {
						if len(psudomoves) > 1 {
							pawnPromotion = true
							promotionMoves = psudomoves
							fmt.Println("1: Queen\n2: Rook\n3: Bishop\n4: Knight")
						} else {
							doMove(board, psudomoves[0])
						}

					} else {
						if newMarked == marked {
							marked = 0
							moves = 0
						} else {
							marked = newMarked
							moves = bitboard.Bitboard(0)
							for _, m := range board.PsudoLegalMoves(false) {
								if marked == m.From {
									moves |= m.To
								}
							}
						}
					}
				}
			}
		}
	} else {
		pressed = false
	}

	if pawnPromotion {
		pawnPromotion = false
		var whitePiece bitboard.PieceType
		var blackPiece bitboard.PieceType

		if ebiten.IsKeyPressed(ebiten.Key1) {
			whitePiece = bitboard.WhiteQueen
			blackPiece = bitboard.BlackQueen
		} else if ebiten.IsKeyPressed(ebiten.Key2) {
			whitePiece = bitboard.WhiteRook
			blackPiece = bitboard.BlackRook
		} else if ebiten.IsKeyPressed(ebiten.Key3) {
			whitePiece = bitboard.WhiteBishop
			blackPiece = bitboard.BlackBishop
		} else if ebiten.IsKeyPressed(ebiten.Key4) {
			whitePiece = bitboard.WhiteKnight
			blackPiece = bitboard.BlackKnight
		} else {
			pawnPromotion = true
		}

		for _, m := range promotionMoves {
			if m.PawnPromotionPiece == whitePiece || m.PawnPromotionPiece == blackPiece {
				doMove(board, m)
			}
		}
	}
}

func doMove(board *bitboard.ChessBoard, m bitboard.Move) {
	temp := *board
	board.DoMove(m)
	if !board.CheckForCheck(board.BlacksTurn) {
		marked = 0
		moves = 0
		m := bitboard.IterativeDeepening(board)
		board.DoMove(m)
	} else {
		*board = temp
		if m.To == marked {
			marked = 0
			moves = 0
		} else {
			marked = m.To
			moves = bitboard.Bitboard(0)
			for _, m := range board.PsudoLegalMoves(false) {
				if marked == m.From {
					moves |= m.To
				}
			}
		}
	}
}
