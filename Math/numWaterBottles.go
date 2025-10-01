//1518. Water Bottles

package math

func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles >= numExchange {
		// fmt.Println(res)
		res += (numBottles / numExchange)
		numBottles = (numBottles / numExchange) + (numBottles % numExchange)
	}
	return res
}
