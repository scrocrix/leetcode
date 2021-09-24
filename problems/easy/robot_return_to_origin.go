package easy

func judgeCircle(moves string) bool {
	// (x,y)
	pos := [2]int{0, 0}
	for _, move := range moves {
		if string(move) == "U" {
			pos[1] = pos[1] + 1
		} else if string(move) == "D" {
			pos[1] = pos[1] - 1
		} else if string(move) == "R" {
			pos[0] = pos[0] + 1
		} else if string(move) == "L" {
			pos[0] = pos[0] - 1
		}
	}
	return pos[0] == 0 && pos[1] == 0
}
