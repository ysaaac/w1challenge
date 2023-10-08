package mario

import "strings"

func PrintPyramid (height int) string {

	var pyramid string
	
	for i := 1; i <= height; i++ {
		spaces := strings.Repeat(" ", height-i)
		hashes := strings.Repeat("#", i)
		line := spaces + hashes + "\n"
		pyramid += line
	}
	return pyramid
}