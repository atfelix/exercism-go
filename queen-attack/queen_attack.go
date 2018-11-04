package queenattack

import (
	"errors"
)

func CanQueenAttack(white, black string) (bool, error) {
	if !isValidPlacement(white, black) {
		return false, errors.New("Invalid placement")
	} 
	return canAttackEachOther(black, white), nil
}

func isValidPlacement(white, black string) bool {
	return isValidSquare(white) && isValidSquare(black) && white != black
}

func isValidSquare(square string) bool {
	return (len(square) == 2 && 
		'a' <= square[0] && square[0] <= 'h' &&
		'1' <= square[1] && square[1] <= '8')
}

func canAttackEachOther(white, black string) bool {
	return (haveSameFile(white, black) ||
				haveSameRank(white, black) ||
				areOnDiagonal(white, black) ||
				areOnOffDiagonal(white, black))
}

func haveSameFile(white, black string) bool {
	return white[0] == black[0]
}

func haveSameRank(white, black string) bool {
	return white[1] == black[1]
}

func areOnDiagonal(white, black string) bool {
	return white[0] - black[0] == white[1] - black[1]
}

func areOnOffDiagonal(white, black string) bool {
	return white[0] - black[0] == black[1] - white[1]
}