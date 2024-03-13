package main

import "fmt"

func main() {
	var qtyIngredient1 int
	var qtyIngredient2 int
	var qtyIngredient3 int

	fmt.Println("Введите количество ингредиента 1:")
	fmt.Scan(&qtyIngredient1)
	fmt.Println("Введите количество ингредиента 2:")
	fmt.Scan(&qtyIngredient2)
	fmt.Println("Введите количество ингредиента 3:")
	fmt.Scan(&qtyIngredient3)

	// Определить цены ингредиентов
	var priceIngredient1 int
	var priceIngredient2 int
	var priceIngredient3 int

	fmt.Println("Введите цену ингредиента 1:")
	fmt.Scan(&priceIngredient1)
	fmt.Println("Введите цену ингредиента 2:")
	fmt.Scan(&priceIngredient2)
	fmt.Println("Введите цену ингредиента 3:")
	fmt.Scan(&priceIngredient3)

	// Расчитать стоимость ингредиентов для одной бутылки зелья
	costIngredient1 := float64(qtyIngredient1) * float64(priceIngredient1)
	costIngredient2 := float64(qtyIngredient2) * float64(priceIngredient2)
	costIngredient3 := float64(qtyIngredient3) * float64(priceIngredient3)

	costIngredients := costIngredient1 + costIngredient2 + costIngredient3

	//Определить цену одной бутылки зелья на рынке
	var pricePotion float64

	fmt.Println("Введите цену одной бутылки зелья на рынке:")
	fmt.Scan(&pricePotion)

	// Преобразовать costIngredients в float
	costIngredientsFloat := float64(costIngredients)
	var discount float64

	fmt.Println("Введите величину скидки в процентах:")
	fmt.Scan(&discount)

	discountFloat := float64(discount) / 100
	withDics := costIngredientsFloat * discountFloat
	OneUnit := ((costIngredientsFloat - withDics) / 5)

	// Расчитать прибыль от одной бутылки зелья
	profit := float64(pricePotion) - OneUnit

	// Напечатать результаты
	fmt.Printf("Стоимость ингредиентов для 5 бутылки зелья: $%.2f\n", costIngredientsFloat)
	fmt.Printf("Цена одной бутылки зелья на рынке: $%.2f\n", pricePotion)
	fmt.Printf("Цена одной бутылки зелье: $%.2f\n", OneUnit)
	fmt.Printf("Прибыль от одной бутылки зелья: $%.2f\n", profit)
}
