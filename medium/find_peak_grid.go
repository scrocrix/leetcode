package medium

import "log"

func FindPeakGrid(mat [][]int) []int {
	r := len(mat) - 1
	c := len(mat[0]) - 1

	peak := []int{}

	for r >= 0 && c >= 0 {
		if len(peak) != 0 {
			break
		}

		// top-left
		if c == 0 && r == 0 {
			if mat[r][c] > mat[r][c+1] && mat[r][c] > mat[r+1][c] {
				peak = []int{r, c}
			}
		}

		// bottom-left
		if r == len(mat)-1 && c == 0 {
			if mat[r][c] > mat[r-1][c] && mat[r][c] > mat[r][c+1] {
				peak = []int{r, c}
			}
		}

		// top-right
		if c == len(mat[0])-1 && r == 0 {
			if mat[r][c] > mat[r][c-1] && mat[r][c] > mat[r+1][c] {
				peak = []int{r, c}
			}
		}

		// bottom-right
		if c == len(mat[0])-1 && r == len(mat)-1 {
			if mat[r][c] > mat[r-1][c] && mat[r][c] > mat[r][c-1] {
				peak = []int{r, c}
			}
		}

		// inside
		if c+1 <= len(mat[0])-1 && c-1 >= 0 {
			log.Println(mat[r][c])
			if r == 0 {
				if mat[r][c] > mat[r+1][c] && mat[r][c] > mat[r][c-1] && mat[r][c] > mat[r][c+1] {
					peak = []int{r, c}
				}
			} else if r == len(mat)-1 {
				if mat[r][c] > mat[r-1][c] && mat[r][c] > mat[r][c-1] && mat[r][c] > mat[r][c+1] {
					peak = []int{r, c}
				}
			} else if mat[r][c] > mat[r-1][c] && mat[r][c] > mat[r+1][c] && mat[r][c] > mat[r][c-1] && mat[r][c] > mat[r][c+1] {
				peak = []int{r, c}
			}
		}

		c = c - 1

		if c == -1 {
			c = len(mat[0]) - 1
			r = r - 1
		}
	}

	return peak
}
