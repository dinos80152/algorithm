package bleedingink

import "math"

type dot struct {
	x        int
	y        int
	darkness int
}

func bleedInk(sheet [][]int, dots []dot) (result [][]int, total int) {
	for _, dot := range dots {
		for y := 0; y < len(sheet); y++ {
			for x := 0; x < len(sheet[y]); x++ {
				bleed := int(float64(dot.darkness) - math.Abs(float64(x-dot.x)) - math.Abs(float64(y-dot.y)))
				if bleed > sheet[y][x] {
					total = total - sheet[y][x] + bleed
					sheet[y][x] = bleed
				}
			}
		}
	}
	return sheet, total
}
