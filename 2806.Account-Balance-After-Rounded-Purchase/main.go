package main

func accountBalanceAfterPurchase(purchaseAmount int) int {
	remainder := purchaseAmount % 10
	if remainder < 5 {
		return 100 + remainder - purchaseAmount
	}
	return 100 - (10-purchaseAmount%10)%10 - purchaseAmount
}
