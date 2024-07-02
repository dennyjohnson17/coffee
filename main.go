package main

import (
	"fmt"
)

func main() {
	var water, milk, coffee, cups int
	var money float64
	water, milk, coffee, cups, money = 500, 500, 500, 500, 100.0

	const waterCostPerMl = 0.001  
	const milkCostPerMl = 0.002    
	const coffeeCostPerGram = 0.04 
	const cupCost = 0.10           

	for {
		var command string
		fmt.Println("Введите команду (buy, fill, take, stats, exit):")
		fmt.Scan(&command)

		switch command {
		case "buy":
			var choice string
			fmt.Println("Выберите напиток (1 - эспрессо, 2 - латте, 3 - капучино):")
			fmt.Scan(&choice)

			switch choice {
			case "1": // эспрессо
				if water >= 250 && coffee >= 16 && cups >= 1 {
					water -= 250
					coffee -= 16
					cups--
					money += 4.0
					fmt.Println("Вы получили эспрессо")
				} else {
					fmt.Println("Недостаточно ресурсов для приготовления эспрессо")
				}
			case "2": // латте
				if water >= 350 && milk >= 75 && coffee >= 20 && cups >= 1 {
					water -= 350
					milk -= 75
					coffee -= 20
					cups--
					money += 7.0
					fmt.Println("Вы получили латте")
				} else {
					fmt.Println("Недостаточно ресурсов для приготовления латте")
				}
			case "3": // капучино
				if water >= 200 && milk >= 100 && coffee >= 12 && cups >= 1 {
					water -= 200
					milk -= 100
					coffee -= 12
					cups--
					money += 6.0
					fmt.Println("Вы получили капучино")
				} else {
					fmt.Println("Недостаточно ресурсов для приготовления капучино")
				}
			default:
				fmt.Println("Неверный выбор")
			}

		case "fill":
			var addWater, addMilk, addCoffee, addCups int

			fmt.Println("Ваши деньги: $", money)

			fmt.Println("Стоимость воды за 100 мл: $", 100*waterCostPerMl)
			fmt.Println("Введите количество воды для добавления (мл):")
			fmt.Scan(&addWater)
			waterCost := float64(addWater) * waterCostPerMl
			fmt.Println("Стоимость воды: $", waterCost)
			if money >= waterCost {
				water += addWater
				money -= waterCost
				fmt.Println("Остаток денег: $", money)
			} else {
				fmt.Println("Недостаточно денег на воду")
			}

			fmt.Println("Стоимость молока за 100 мл: $", 100*milkCostPerMl)
			fmt.Println("Введите количество молока для добавления (мл):")
			fmt.Scan(&addMilk)
			milkCost := float64(addMilk) * milkCostPerMl
			fmt.Println("Стоимость молока: $", milkCost)
			if money >= milkCost {
				milk += addMilk
				money -= milkCost
				fmt.Println("Остаток денег: $", money)
			} else {
				fmt.Println("Недостаточно денег на молоко")
			}

			fmt.Println("Стоимость кофе за 100 г: $", 100*coffeeCostPerGram)
			fmt.Println("Введите количество кофе для добавления (г):")
			fmt.Scan(&addCoffee)
			coffeeCost := float64(addCoffee) * coffeeCostPerGram
			fmt.Println("Стоимость кофе: $", coffeeCost)
			if money >= coffeeCost {
				coffee += addCoffee
				money -= coffeeCost
				fmt.Println("Остаток денег: $", money)
			} else {
				fmt.Println("Недостаточно денег на кофе")
			}

			fmt.Println("Введите количество стаканчиков для добавления (шт):")
			fmt.Scan(&addCups)
			cupCostTotal := float64(addCups) * cupCost
			fmt.Println("Стоимость стаканчиков: $", cupCostTotal)
			if money >= cupCostTotal {
				cups += addCups
				money -= cupCostTotal
				fmt.Println("Остаток денег: $", money)
			} else {
				fmt.Println("Недостаточно денег на стаканчики")
			}

		case "take":
			fmt.Println("Вы забрали $", money)
			money = 0.0

		case "stats":
			fmt.Println("Вода:", water, "мл")
			fmt.Println("Молоко:", milk, "мл")
			fmt.Println("Кофе:", coffee, "г")
			fmt.Println("Стаканчики:", cups, "шт")
			fmt.Println("Деньги: $", money)

		case "exit":
			return

		default:
			fmt.Println("Неверная команда")
		}
	}
}
