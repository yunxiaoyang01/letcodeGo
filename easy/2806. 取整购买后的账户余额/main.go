package main

func main() {

}
func accountBalanceAfterPurchase(purchaseAmount int) int {
	r := purchaseAmount % 10
	if r < 5 {
		purchaseAmount -= r
	} else {
		purchaseAmount += 10 - r
	}
	return 100 - purchaseAmount
}
