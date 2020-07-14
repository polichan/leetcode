package main

func main() {

}

func calculateMinimumHP(dungeon [][]int) int {
	row := len(dungeon)
	col := len(dungeon[0])

	hp := make([][]int, row + 1)
	for i := range hp {
		hp[i] = make([]int, col + 1)
	}

	for a:=0; a<len(hp); a++ {
		for b := 0; b<len(hp[0]); b++ {
			hp[a][b] = 99999
		}
	}

	hp[row][col - 1]  = 1
	hp[row - 1][col] = 1
	for y := row -1; y >= 0; y-- {
		for x:=col -1; x >= 0 ; x-- {
			hp[y][x] = Max(1, Min(hp[y+1][x], hp[y][x+1]) - dungeon[y][x])
		}
	}
	return hp[0][0]
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int)int  {
	if x>y {
		return x
	}
	return y
}




