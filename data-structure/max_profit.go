package data_structure

/*
Best Time to Buy and Sell Stock
配列priceが与えられ、price[i]は指定された銘柄のi日目の価格である。
あなたは、ある銘柄を買うために1日を選び、その銘柄を売るために将来の異なる日を選ぶことによって、
利益を最大化したいと思います。
この取引で達成できる最大の利益を返しなさい。もし利益が出なければ、0を返せ。

Input: prices = [7,1,5,3,6,4]
Output: 5
2日目（価格=1）に買い、5日目（価格=6）に売り、利益=6-1=5。
2日目に買って1日目に売ることは、売る前に買わなければならないので、許されないことに注意
*/

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min, maxSale := prices[0], 0
	for _, price := range prices {
		if price < min {
			min = price
		} else if (price - min) > maxSale {
			maxSale = price - min
		}
	}
	return maxSale
}
