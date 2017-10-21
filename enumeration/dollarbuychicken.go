package enumeration

// 公雞 5 元，母雞 3 元，小雞 3 個 1 元
// 給多少錢與總共想買幾隻雞，問購買組合？
func dollarBuyChicken(dollar, quantity int) (ans [][3]int) {
	var xPrice, yPrice = 5, 3
	for x := 0; x <= dollar/xPrice; x++ {
		restDollar := dollar - x*xPrice
		for y := 0; y <= restDollar/yPrice; y++ {
			z := quantity - x - y
			if z%3 == 0 {
				if (5*x + 3*y + z/3) == dollar {
					ans = append(ans, [3]int{x, y, z})
				}
			}
		}
	}
	return ans
}
