package q1

package main

import (
	"errors"
	"fmt"
)

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase <= 0 {
		return 0, errors.New("valor da compra inválido")
	}

	var totalCompras float64
	for _, compra := range purchaseHistory {
		totalCompras += compra
	}

	mediaCompras := totalCompras / float64(len(purchaseHistory))

	var discount float64

	if totalCompras > 1000 {
		discount = currentPurchase * 0.10
	} else if totalCompras <= 1000 && totalCompras > 500 {
		discount = currentPurchase * 0.05
	} else if totalCompras <= 500 {
		discount = currentPurchase * 0.02
	}

	if len(purchaseHistory) == 0 {
		discount = currentPurchase * 0.10
	}

	if mediaCompras > 1000 && discount < (currentPurchase*0.20) {
		discount = currentPurchase * 0.20
	}

	return discount, nil
}
