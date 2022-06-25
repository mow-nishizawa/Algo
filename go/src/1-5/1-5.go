package main

import "fmt"

func main() {
	const money = 1000
	const coinNum = 15
	const fiveHund = 500
	const oneHund = 100
	const fifty = 50
	const ten = 10
	var total = 0

	for i := money / fiveHund; i >= 0; i-- {
		if i > coinNum {
			break
		}
		var leftMoney = money - (fiveHund * i)
		if leftMoney == 0 {
			total++
			fmt.Printf("500: %d\n", i)
			continue
		}
		for j := leftMoney / oneHund; j >= 0; j-- {
			if i+j > coinNum {
				break
			}
			var leftMoneyOneHund = leftMoney - (oneHund * j)
			if leftMoneyOneHund == 0 {
				total++
				fmt.Printf("500: %d, 100: %d\n", i, j)
				continue
			}
			for k := leftMoneyOneHund / fifty; k >= 0; k-- {
				if i+j+k > coinNum {
					break
				}
				var leftMoneyFifty = leftMoneyOneHund - (fifty * k)
				if leftMoneyFifty == 0 {
					total++
					fmt.Printf("500: %d, 100: %d, 50: %d\n", i, j, k)
					continue
				}
				for l := leftMoneyFifty / ten; l >= 0; l-- {
					if i+j+k+l > coinNum {
						break
					}
					var leftMoneyTen = leftMoneyFifty - (ten * l)
					if leftMoneyTen == 0 {
						total++
						fmt.Printf("500: %d, 100: %d, 50: %d, 10: %d\n", i, j, k, l)
						continue
					}
				}
			}
		}
	}
	fmt.Printf("total: %d\n", total)
}
