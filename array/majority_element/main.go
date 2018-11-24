// TODO: best solution
package main

func main() {

}

func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	var most int
	for _, num := range nums {
		countMap[num]++
		if countMap[num] > countMap[most] {
			most = num
		}
	}
	return most
}
