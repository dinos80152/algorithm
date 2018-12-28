/*
Imagine we have an image. We'll represent this image as a simple 2D array where every pixel is a 1 or a 0. The image you get is known to have a single rectangle of 0s on a background of 1s.

Write a function that takes in the image and returns the coordinates of the rectangle of 0's -- either top-left and bottom-right; or top-left, width, and height.

Sample output:
x: 3, y: 2, width: 3, height: 2
2,3 3,5
3,2 5,3 -- it's ok to reverse columns/rows as long as you're consistent
*/

package twodarray

import "fmt"

func main() {
	image := [][]int{
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 0, 0, 0, 1},
		[]int{1, 1, 1, 0, 0, 0, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
	}
	image2 := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0},
	}
	image3 := [][]int{
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
	}
	coordinate(image)
	coordinate(image2)
	fmt.Println(coordinate(image3))
}

func coordinate(image [][]int) ([]int, int) {
	var coordinates [][]int
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] == 0 {
				coordinates = append(coordinates, []int{j, i})
			}
		}
	}

	if len(coordinates) == 0 {
		return []int{}, 0
	}

	//width := coordinates[len(coordinates)-1][0] - coordinates[0][0] + 1
	height := coordinates[len(coordinates)-1][1] - coordinates[0][1] + 1
	return coordinates[0], height
}
